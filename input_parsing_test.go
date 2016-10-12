package console_test

import (
	"testing"

	"github.com/eidolon/console"
	"github.com/eidolon/console/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("should return an empty Input if no parameters are given", func(t *testing.T) {
		params := []string{}
		input := console.ParseInput(params)

		assert.True(t, len(input.Arguments) == 0, "Expected length to be 0")
		assert.True(t, len(input.Options) == 0, "Expected length to be 0")
	})

	t.Run("should parse strings not starting with '-' or '--' as arguments", func(t *testing.T) {
		params := []string{
			"hello",
			"world",
		}

		input := console.ParseInput(params)

		assert.True(t, len(input.Arguments) == 2, "Expected length to be 2")
		assert.True(t, len(input.Options) == 0, "Expected length to be 0")
	})

	t.Run("should retain argument order", func(t *testing.T) {
		t.Skip()
	})

	t.Run("should parse '--' as an argument", func(t *testing.T) {
		t.Skip()
	})
}
