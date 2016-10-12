package console

// Output abstracts application output. This is mainly useful for testing, as a different writer can
// be passed to capture output in an easy to test manner.
type Output struct{}

// NewOutput creates a new Output.
func NewOutput() *Output {
	return &Output{}
}

// @todo: Output (low priority):
// @todo: - Allow specifying a writer to write to.
// @todo: - Add Write, for writing strings to a the given writer.
