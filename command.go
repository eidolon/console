package console

// ConfigureFunc is a function to mutate the input definition to add arguments and options.
type ConfigureFunc func(*Definition)

// ExecuteFunc is a function to perform whatever task this command does.
type ExecuteFunc func() error

// Command represents a command to run in an application.
type Command struct {
	// The name of the command.
	Name string
	// The description of the command.
	Description string
	// Help message for the command.
	Help string
	// Function to configure command-level parameters.
	Configure ConfigureFunc
	// Function to execute when this command is requested.
	Execute ExecuteFunc
}
