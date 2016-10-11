package console

// Option value modes.
const (
	OptionValueNone OptionValueMode = iota - 1
	OptionValueOptional
	OptionValueRequired
)

// OptionValueMode represents the different potential requirements of an option's value.
type OptionValueMode int

// Option provides the internal representation of an input option paremeter.
type Option struct {
	// The names of this option.
	Names []string
	// The description of this option.
	Description string
	// The value that this option references.
	Value ParameterValue
	// Does this option take a value? Is it optional, or required?
	ValueMode OptionValueMode
	// The name of the value (shown in contextual help).
	ValueName string
}
