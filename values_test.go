package console

import (
	"testing"
	"time"
)

func TestBoolValue(t *testing.T) {
	t.Run("BoolValue()", func(t *testing.T) {
		t.Run("should create a boolValue with the given value", func(t *testing.T) {
			truthy := true
			falsey := false

			truthyValue := BoolValue(&truthy)
			falseyValue := BoolValue(&falsey)

			if actual := truthyValue.String(); actual != "true" {
				t.Errorf("Expected 'true', got '%s'.", actual)
			}

			if actual := falseyValue.String(); actual != "false" {
				t.Errorf("Expected 'true', got '%s'.", actual)
			}
		})
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value boolValue

			valid := []string{
				"1",
				"t",
				"T",
				"TRUE",
				"true",
				"True",
				"0",
				"f",
				"F",
				"FALSE",
				"false",
				"False",
			}

			for _, item := range valid {
				err := value.Set(item)
				if err != nil {
					t.Errorf("Unexpected error: '%s'.", err)
				}
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value boolValue

			invalid := []string{
				"",
				"yes",
				"no",
				"Y",
				"N",
			}

			for _, item := range invalid {
				err := value.Set(item)
				if err == nil {
					t.Error("Expected error, did not receive one.")
				}
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := true
			value := BoolValue(&ref)

			if ref != true {
				t.Errorf("Expected true, got %v", ref)
			}

			value.Set("false")

			if ref != false {
				t.Errorf("Expected false, got %v", ref)
			}

			value.Set("true")

			if ref != true {
				t.Errorf("Expected true, got %v", ref)
			}
		})
	})

	t.Run("String()", func(t *testing.T) {
		t.Run("should return either 'true' or 'false'", func(t *testing.T) {
			inOut := map[string]string{
				"1":     "true",
				"t":     "true",
				"T":     "true",
				"TRUE":  "true",
				"true":  "true",
				"True":  "true",
				"0":     "false",
				"f":     "false",
				"F":     "false",
				"FALSE": "false",
				"false": "false",
				"False": "false",
			}

			for in, expected := range inOut {
				value := new(boolValue)
				value.Set(in)

				actual := value.String()

				if actual != expected {
					t.Errorf("Expected '%s', got '%s'.", expected, actual)
				}
			}
		})
	})

	t.Run("FlagValue()", func(t *testing.T) {
		t.Run("should always return 'true'", func(t *testing.T) {
			var value boolValue

			expected := "true"
			actual := value.FlagValue()

			if actual != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, actual)
			}
		})
	})
}

func TestDurationValue(t *testing.T) {
	t.Run("DurationValue()", func(t *testing.T) {
		t.Run("should create a duractionValue with the given value", func(t *testing.T) {
			duration := time.Second
			durationValue := DurationValue(&duration)

			if actual := durationValue.String(); actual != "1s" {
				t.Errorf("Expected '1s', got '%s'.", actual)
			}
		})
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid duration values", func(t *testing.T) {
			var value durationValue

			valid := []string{
				"5us",
				"999ns",
				"1s",
				"3m",
				"1h33m2s",
				"5h",
				"365h",
			}

			for _, item := range valid {
				err := value.Set(item)
				if err != nil {
					t.Errorf("Unexpected error: '%s'.", err)
				}
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value durationValue

			invalid := []string{
				"",
				"1d",
				"20y",
				"20 decades",
			}

			for _, item := range invalid {
				err := value.Set(item)
				if err == nil {
					t.Error("Expected error, did not receive one.")
				}
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := time.Second
			value := DurationValue(&ref)

			if ref != time.Second {
				t.Errorf("Expected true, got %v", ref)
			}

			value.Set("1m")

			if ref != time.Minute {
				t.Errorf("Expected false, got %v", ref)
			}

			value.Set("1h")

			if ref != time.Hour {
				t.Errorf("Expected true, got %v", ref)
			}
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"1us": "1µs",
			"1ns": "1ns",
			"1s":  "1s",
			"1m":  "1m0s",
			"1h":  "1h0m0s",
			"5s":  "5s",
			"10h": "10h0m0s",
		}

		for in, expected := range inOut {
			value := new(durationValue)
			value.Set(in)

			actual := value.String()

			if actual != expected {
				t.Errorf("Expected '%s', got '%s'.", expected, actual)
			}
		}
	})
}
