package console

import (
	"fmt"
	"strings"
)

// ParseInput takes an array of strings (typically arguments to the application), and parses them
// into the raw Input type.
func ParseInput(params []string) *Input {
	var result Input

	processOptions := true

	for _, param := range params {
		if param == "--" {
			processOptions = false
			continue
		}

		paramLen := len(param)
		isLongOpt := paramLen > 2 && strings.HasPrefix(param, "--")
		isShortOpt := paramLen > 1 && strings.HasPrefix(param, "-")

		if processOptions && isLongOpt {
			result.Options = append(result.Options, parseOption(param, "--")...)
		} else if processOptions && isShortOpt {
			result.Options = append(result.Options, parseOption(param, "-")...)
		} else {
			result.Arguments = append(result.Arguments, InputArgument{Value: param})
		}
	}

	return &result
}

// parseOption parses an input option with the given prefix (e.g. '-', or '--'). It returns an array
// because short options can contain multiple options without values.
func parseOption(option string, prefix string) []InputOption {
	var results []InputOption

	trimmed := strings.TrimPrefix(option, prefix)
	split := strings.SplitN(trimmed, "=", 2)

	var key string
	var val string

	if len(split) >= 1 {
		key = split[0]
	}

	if len(split) == 2 {
		val = split[1]
	}

	if prefix == "-" {
		results = append(results, parseShortOption(key, val)...)
	}

	if prefix == "--" {
		results = append(results, parseLongOption(key, val))
	}

	return results
}

func parseLongOption(key, value string) InputOption {
	return InputOption{Name: key, Value: value}
}

func parseShortOption(key, value string) []InputOption {
	var results []InputOption

	// Convert key into rune slice, so we can iterate over each rune properly.
	runes := []rune(key)

	// Folded options
	if len(key) > 1 {
		// We want to handle the last run differently, so that the value following all of the
		// options can be given to the last option.
		for i := 0; i < len(runes)-1; i++ {
			results = append(results, InputOption{Name: fmt.Sprintf("%c", runes[i])})
		}

		results = append(results, InputOption{Name: fmt.Sprintf("%c", runes[len(runes)-1]), Value: value})
	} else {
		results = append(results, InputOption{Name: key, Value: value})
	}

	return results
}
