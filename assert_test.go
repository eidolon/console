package console

import (
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// assert asserts that the given condition is truthy.
func assert(t *testing.T, condition bool, message string) {
	_, file, line, _ := runtime.Caller(1)

	if !condition {
		t.Errorf(
			"assert: %s:%d: %s",
			trimLocation(file),
			line,
			message,
		)
	}
}

// assertEquals asserts that the given expected and actual arguments are equal.
func assertEqual(t *testing.T, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(1)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"assert: %s:%d: Expected %v (type %v) to equal %v (type %v).",
			trimLocation(file),
			line,
			actual,
			reflect.TypeOf(actual),
			expected,
			reflect.TypeOf(expected),
		)
	}
}

// assertUnequal asserts that the given expected and actual arguments are not equal.
func assertUnequal(t *testing.T, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(1)

	if reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"assert: %s:%d: Expected %v (type %v) not to equal got %v (type %v).",
			trimLocation(file),
			line,
			actual,
			reflect.TypeOf(actual),
			expected,
			reflect.TypeOf(expected),
		)
	}
}

// assertOK checks that an error was not produced, and reacts accordingly.
func assertOK(t *testing.T, err error) {
	_, file, line, _ := runtime.Caller(1)

	if err != nil {
		t.Errorf(
			"assert: %s:%d: Unexpected error: '%s'.",
			trimLocation(file),
			line,
			err.Error(),
		)
	}
}

// assertNotOK checks that an error was produced, and reacts accordingly.
func assertNotOK(t *testing.T, err error) {
	_, file, line, _ := runtime.Caller(1)

	if err == nil {
		t.Errorf(
			"assert: %s:%d: Expected error, but none was given.",
			trimLocation(file),
			line,
		)
	}
}

// trimLocation takes an absolute file path, and returns a much shorter relative path.
func trimLocation(file string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return strings.Replace(file, cwd+"/", "", -1)
}
