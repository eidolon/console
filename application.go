package console

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

	// Show help?
	help bool
	// Application input.
	input Input
}

// NewApplication creates a new Application with some sane defaults.
func NewApplication(name string, version string) *Application {
	return &Application{
		Name:      name,
		UsageName: filepath.Base(os.Args[0]),
		Version:   version,
		Writer:    os.Stderr,
	}
}

// Quick check to see if a help flag is set, ignoring values. Uses raw input, not mapped input.
func (a *Application) hasHelpOption() bool {
	for _, opt := range a.input.Options {
		if opt.Name == "help" {
			return true
		}
	}

	return false
}

// Run the configured application, with the given input.
func (a *Application) Run(params []string) int {
	// Create input and output.
	input := ParseInput(params)
	output := Output{}
	definition := NewDefinition()

	// Assign input to application.
	a.input = input

	command := a.findCommandInInput()
	if command != nil && command.Configure != nil {
		command.Configure(definition)
	}

	noCmdExecute := command == nil || command.Execute == nil

	if noCmdExecute {
		// @todo: Use output? Check if help should be shown?
		fmt.Fprintln(a.Writer, "NYI")
		return 2
	}

	err := MapInput(*definition, input)
	if err != nil {
		fmt.Fprintln(a.Writer, err)
		fmt.Fprintln(a.Writer, fmt.Sprintf("Try '%s --help' for more information.", a.UsageName))
		return 1
	}

	err = command.Execute(input, output)
	if err != nil {
		fmt.Fprintln(a.Writer, err)
		fmt.Fprintln(
			a.Writer,
			fmt.Sprintf(
				"Try '%s %s --help' for more information.",
				a.UsageName,
				command.Name,
			),
		)
		return 1
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
