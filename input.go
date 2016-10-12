package console

// Input represents the raw application input, in a slightly more organised way, and provides
// helpers for retrieving that information. This allows commands to get application-wide input,
// instead of only the command-specific input.
type Input struct {
	Arguments []InputArgument
	Options   []InputOption
}

// InputArgument represents the raw data parsed as arguments, really this is just the value.
type InputArgument struct {
	Value string
}

// InputOption represents the raw data parsed as options. This includes it's name and it's value.
// The value can be "" if no value is given (i.e. if the option is a flag).
type InputOption struct {
	Name  string
	Value string
}

// @todo: Input (low priority):
// @todo: - Add method for retrieving argument by index.
// @todo: - Add method for retrieving option by name.
// @todo: - Add method for retrieving option by names (first matching).
