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
	// Array of commands that can be run. May contain sub-commands.
	Commands []Command
	// Function to configure application-level parameters, realistically should just be options.
	Configure ConfigureFunc
	// Writer to write output to.
	Writer io.Writer

	// Application input.
	input *Input
}

// NewApplication creates a new Application with some sane defaults.
func NewApplication(name string, version string) *Application {
	return &Application{
		Name:      name,
		UsageName: filepath.Base(os.Args[0]),
		Version:   version,
		Writer:    os.Stdout,
	}
}

// Run runs the configured application, with the given input.
func (a *Application) Run(params []string) int {
	// Create input and output.
	input := ParseInput(params)
	output := NewOutput(a.Writer)
	definition := NewDefinition()

	// Assign input to application.
	a.input = input
	a.preConfigure(definition)

	if a.Configure != nil {
		// Similar to a.preConfigure, this is useful for customising application-level help text.
		// This however will work just like a command's Configure function.
		a.Configure(definition)

		// Throw away any defined arguments, as they cannot be used, and will disrupt help text.
		definition.arguments = make(map[string]parameters.Argument)
	}

	command := a.findCommandInInput()
	if command != nil && command.Configure != nil {
		command.Configure(definition)
	}

	if a.hasHelpOption() || (command == nil || command.Execute == nil) {
		a.showHelp(output, command)
		return 100
	}

	err := MapInput(definition, input)
	if err != nil {
		output.Println(err)
		output.Printf("Try '%s --help' for more information.\n", a.UsageName)
		return 101
	}

	err = command.Execute(input, output)
	if err != nil {
		output.Println(err)
		output.Printf("Try '%s %s --help' for more information.\n", a.UsageName, command.Name)
		return 102
	}

	return 0
}

// AddCommands adds commands to the application.
func (a *Application) AddCommands(commands []Command) {
	a.Commands = append(a.Commands, commands...)
}

// AddCommand adds a command to the application.
func (a *Application) AddCommand(command Command) {
	a.Commands = append(a.Commands, command)
}

// findCommand attempts to find the command to run based on the raw input.
func (a *Application) findCommandInInput() *Command {
	// There can't be a command if there are no arguments!
	if len(a.input.Arguments) == 0 {
		return nil
	}

	var command *Command
	for i, cmd := range a.Commands {
		if cmd.Name == a.input.Arguments[0].Value {
			command = &a.Commands[i]
			break
		}
	}

	if command != nil {
		// Forget about the command name, the command that will be run shouldn't use it as input
		a.input.Arguments = a.input.Arguments[1:]

		return command
	}

	return nil
}

// hasHelpOption checks to see if a help flag is set, ignoring values. Uses raw input, not mapped
// input.
func (a *Application) hasHelpOption() bool {
	for _, opt := range a.input.Options {
		if opt.Name == "help" || opt.Name == "h" {
			return true
		}
	}

	return false
}

// preConfigure configures pre-defined parameters. This is solely defined for help output.
func (a *Application) preConfigure(definition *Definition) {
	var help bool
	definition.AddOption(parameters.NewBoolValue(&help), "-h, --help", "Display contextual help?")
}

// showHelp shows contextual help.
func (a *Application) showHelp(output *Output, command *Command) {
	if command != nil {
		output.Println(DescribeCommand(a, command))
	} else {
		output.Println(DescribeApplication(a))
	}
}
