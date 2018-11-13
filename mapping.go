package console

import (
	"strings"

	"github.com/eidolon/console/parameters"
)

// Mapping and parsing when combined check all input provided as application arguments and
// environment variables against all input defined in a Definition.

// MapInput2 iterates over all defined possible input in the definition, and attempts to set the
// value defined in the input, or the environment. In the process of mapping, the input will also be
// validated (i.e. missing required params or values will be identified, and values of the wrong
// type in the input or env will be identified).
func MapInput2(definition *Definition, input Input, env []string) error {
	return nil
}

// ParseInput2 takes the raw input, and regardless of what is actually defined in the definition,
// categorising the input as either arguments or options. In other words, the raw input is iterated
// over, not the definition's parameters. The definition is used so that we can identify options
// that should have values and consume the next argument as it's value.
func ParseInput2(definition *Definition, args []string) *Input {
	var input Input
	var optsEnded bool

	// We don't range, because we can modify `i` in the middle of the loop this way. This allows us
	// to consume the next argument if we want (and if it's available).
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			optsEnded = true
			continue
		}

		argLen := len(arg)
		isLongOpt := argLen > 2 && strings.HasPrefix(arg, "--")
		isShortOpt := argLen > 1 && strings.HasPrefix(arg, "-")

		if !optsEnded && (isLongOpt || isShortOpt) {
			var options []InputOption

			if isLongOpt {
				options = parseOption(arg, "--")
			} else {
				options = parseOption(arg, "-")
			}

			// mappedOptions is a temporary place for the options parsed from this one argument to
			// be stored. We need this so we can always identify if we actually parsed any options
			// at all, and so we can get the last option for this argument all of the time.
			var mappedOptions []InputOption

			for _, option := range options {
				// All options will be mapped, regardless. Only the last option will have any value.
				mappedOptions = append(mappedOptions, option)
			}

			mappedOptionsLen := len(mappedOptions)

			if mappedOptionsLen > 0 {
				lastOption := mappedOptions[mappedOptionsLen-1]

				defOpt, exists := definition.options[lastOption.Name]
				if !exists {
					// We don't care about options that don't exist in the definition. We shouldn't
					// be consuming arguments for them, because they won't require a value.
					break
				}

				isRequired := defOpt.ValueMode == parameters.OptionValueRequired
				hasArgsLeft := len(args) > (i + 1) // Length required for next is +2, not +1.
				hasNoValYet := lastOption.Value == ""

				// If the value is required, but we don't yet have a value on the option, this means
				// we'll consume the next argument following the option and treat it as the value.
				if isRequired && hasNoValYet && hasArgsLeft {
					lastOption.Value = args[i+1]
					i++
				}

				mappedOptions[mappedOptionsLen-1] = lastOption
			}

			input.Options = append(input.Options, mappedOptions...)
		} else {
			input.Arguments = append(input.Arguments, InputArgument{Value: arg})
		}
	}

	return &input
}
