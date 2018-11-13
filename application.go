package console

import (
	"io"
	"os"
	"path/filepath"

	"github.com/eidolon/console/parameters"
)

// Application represents the heart of the console application. It is what orchestrates running
// commands, initiates input parsing, mapping, and validation; and will handle failure for each of
// those tasks.
type Application struct {
	// The name of the application.
	Name string
	// The name of the application printed in usage information, defaults to the binary's filename.
	UsageName string
	// The version of the application.
	Version string
	// Application logo, shown in help output
	Logo string
	// Help message for the application.
	Help string
	// Writer to write output to.
	Writer io.Writer

	// Slice of commands that can be run. May contain sub-commands.
	commands []*Command
	// Slice of global options.
	globalOptionDefinitions []OptionDefinition
	// Application definition
	definition *Definition
	// Application input.
	input *Input
	// Application output.
	output *Output
	// The path taken to reach the current command (used for help text).
	path []string
}

// NewApplication creates a new Application with some sane defaults.
func NewApplication(name string, version string) *Application {
	return &Application{
		Name:       name,
		UsageName:  filepath.Base(os.Args[0]),
		Version:    version,
		Writer:     os.Stdout,
		definition: NewDefinition(),
	}
}

// Run runs the configured application, with the given input.
func (a *Application) Run(argv []string, env []string) int {
	if a.definition == nil {
		panic("attempted to start application with nil definition")
	}

	// Set output at runtime, so that it's available for everything else that could use it, and
	// up-to-date with what the user has requested their io.Writer to be.
	a.output = NewOutput(a.Writer)

	a.configure(a.definition)

	// @TODO: Could we handle global options before we do anything with commands? It wouldn't be too
	// useful for the `help` argument because we need to know the context (i.e. cmd) we're
	// running to show the right thing.
	cmd, path := a.resolveCommand(argv)
	if cmd != nil && cmd.Configure != nil {
		cmd.Configure(a.definition)
	}

	// Trim argv so that the command path is not left in and sent to commands.
	argv = argv[len(path):]

	if a.hasHelpOption(argv) || (cmd == nil || cmd.Execute == nil) {
		a.showHelp(cmd, path)
		return 100
	}

	// Assign input to application.
	a.input = ParseInput2(a.definition, argv)

	err := MapInput(a.definition, a.input, env)
	if err != nil {
		a.output.Println(err)
		a.output.Printf("Try '%s --help' for more information.\n", a.UsageName)
		return 101
	}

	err = cmd.Execute(a.input, a.output)
	if err != nil {
		a.output.Println(err)
		a.output.Printf("Try '%s %s --help' for more information.\n", a.UsageName, cmd.Name)
		return 1
	}

	return a.output.exitCode
}

// AddCommands adds commands to the application.
func (a *Application) AddCommands(commands []*Command) {
	a.commands = append(a.commands, commands...)
}

// AddCommand adds a command to the application.
func (a *Application) AddCommand(command *Command) {
	a.commands = append(a.commands, command)
}

// Commands gets the sub-commands on an application.
func (a *Application) Commands() []*Command {
	return a.commands
}

func (a *Application) AddGlobalOption(definition OptionDefinition) {
	a.globalOptionDefinitions = append(a.globalOptionDefinitions, definition)
}

// resolveCommand attempts to find the command to run based on the raw input.
func (a *Application) resolveCommand(args []string) (*Command, []string) {
	var loop func(depth int, container CommandContainer) *Command
	var path []string

	loop = func(depth int, container CommandContainer) *Command {
		if len(args) < (depth + 1) {
			return nil
		}

		var command *Command
		for _, cmd := range container.Commands() {
			isNameMatch := cmd.Name == args[depth]
			isAliasMatch := cmd.Alias == args[depth]

			if isNameMatch || isAliasMatch {
				command = cmd
				// Add to breadcrumb trail...
				path = append(path, cmd.Name)
				break
			}
		}

		if command != nil {
			subCommand := loop(depth+1, command)

			if subCommand != nil {
				command = subCommand
			}
		}

		return command
	}

	return loop(0, a), path
}

// hasHelpOption checks to see if a help flag is set, ignoring values. Uses raw args sent to the
// application.
func (a *Application) hasHelpOption(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}

	return false
}

// configure configures pre-defined parameters. This is solely defined for help output.
func (a *Application) configure(definition *Definition) {
	var help bool

	definition.AddOption(OptionDefinition{
		Value: parameters.NewBoolValue(&help),
		Spec:  "-h, --help",
		Desc:  "Display contextual help?",
	})

	for _, opt := range a.globalOptionDefinitions {
		definition.AddOption(opt)
	}
}

// showHelp shows contextual help.
func (a *Application) showHelp(command *Command, path []string) {
	if command != nil {
		a.output.Println(DescribeCommand(a, command, path))
	} else {
		a.output.Println(DescribeApplication(a))
	}
}
