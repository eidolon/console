package console

import "testing"

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
