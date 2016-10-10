package console

const (
	optionValueNone = iota - 1
	optionValueOptional
	optionValueRequired
)

type optionValueMode int

// Option provides the internal representation of an input option paremeter.
type Option struct {
	// The names of this option.
	Names []string
	// The description of this option.
	Description string
	// The value that this option references.
	Value ParameterValue
	// Does this option take a value? Is it optional, or required?
	ValueMode optionValueMode
	// The name of the value (shown in contextual help).
	ValueName string
}
