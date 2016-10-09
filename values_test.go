package console

import (
	"net/url"
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

			assertEqual(t, truthyValue.String(), "true")
			assertEqual(t, falseyValue.String(), "false")
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
				assertOK(t, err)
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
				assertNotOK(t, err)
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := true
			value := BoolValue(&ref)

			assertEqual(t, ref, true)

			value.Set("false")
			assertEqual(t, ref, false)

			value.Set("true")
			assertEqual(t, ref, true)
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
				assertEqual(t, expected, actual)
			}
		})
	})

	t.Run("FlagValue()", func(t *testing.T) {
		t.Run("should always return 'true'", func(t *testing.T) {
			var value boolValue

			assertEqual(t, "true", value.FlagValue())
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
				assertOK(t, err)
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
				assertNotOK(t, err)
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := time.Second
			value := DurationValue(&ref)

			assertEqual(t, ref, time.Second)

			value.Set("1m")
			assertEqual(t, ref, time.Minute)

			value.Set("1h")
			assertEqual(t, ref, time.Hour)
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
			assertEqual(t, expected, actual)
		}
	})
}

func TestFloat32Value(t *testing.T) {

}

func TestFloat64Value(t *testing.T) {

}

func TestIntValue(t *testing.T) {

}

func TestIPValue(t *testing.T) {

}

func TestStringValue(t *testing.T) {

}

func TestUrlValue(t *testing.T) {
	t.Run("Set()", func(t *testing.T) {
		t.Run("should modify the URL that it references", func(t *testing.T) {
			oldUrl := "https://www.google.co.uk/"
			newUrl := "https://www.elliotdwright.com/"

			ref, err := url.Parse(oldUrl)
			assertOK(t, err)

			value := URLValue(ref)
			assertEqual(t, oldUrl, ref.String())

			value.Set(newUrl)
			assertEqual(t, newUrl, ref.String())
		})
	})
}
