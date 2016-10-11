package console

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/eidolon/console/assert"
)

func TestBoolValue(t *testing.T) {
	t.Run("newBoolValue()", func(t *testing.T) {
		truthy := true
		falsey := false

		truthyValue := newBoolValue(&truthy)
		falseyValue := newBoolValue(&falsey)

		assert.Equal(t, truthyValue.String(), "true")
		assert.Equal(t, falseyValue.String(), "false")
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
				assert.OK(t, err)
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
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := true
			value := newBoolValue(&ref)

			assert.Equal(t, true, ref)

			value.Set("false")
			assert.Equal(t, false, ref)

			value.Set("true")
			assert.Equal(t, true, ref)
		})
	})

	t.Run("String()", func(t *testing.T) {
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
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("FlagValue()", func(t *testing.T) {
		var value boolValue

		assert.Equal(t, "true", value.FlagValue())
	})
}

func TestDurationValue(t *testing.T) {
	t.Run("newDurationValue()", func(t *testing.T) {
		duration := time.Second
		durationValue := newDurationValue(&duration)

		assert.Equal(t, "1s", durationValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
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
				assert.OK(t, err)
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
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the bool that it references", func(t *testing.T) {
			ref := time.Second
			value := newDurationValue(&ref)

			assert.Equal(t, time.Second, ref)

			value.Set("1m")
			assert.Equal(t, time.Minute, ref)

			value.Set("1h")
			assert.Equal(t, ref, time.Hour)
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
			assert.Equal(t, expected, actual)
		}
	})
}

func TestFloat32Value(t *testing.T) {
	t.Run("newFloat32Value()", func(t *testing.T) {
		float := float32(3.14)
		floatValue := newFloat32Value(&float)

		assert.Equal(t, "3.14", floatValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value float32Value

			valid := []string{
				"3",
				"3.1",
				"3.14",
				"314.159e-2",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value float32Value

			invalid := []string{
				"",
				"Hello, World!",
				"Three point one four",
			}

			for _, item := range invalid {
				err := value.Set(item)
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the float32 that it references", func(t *testing.T) {
			ref := float32(3.14)
			value := newFloat32Value(&ref)

			assert.Equal(t, float32(3.14), ref)

			value.Set("3.14159")
			assert.Equal(t, float32(3.14159), ref)

			value.Set("10")
			assert.Equal(t, float32(10), ref)
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"3":          "3",
			"3.14":       "3.14",
			"3.14159":    "3.14159",
			"314.159e-2": "3.14159",
		}

		for in, expected := range inOut {
			value := new(float32Value)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}

func TestFloat64Value(t *testing.T) {
	t.Run("newFloat64Value()", func(t *testing.T) {
		float := float64(3.14)
		floatValue := newFloat64Value(&float)

		assert.Equal(t, "3.14", floatValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value float64Value

			valid := []string{
				"3",
				"3.1",
				"3.14",
				"314.159e-2",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value float64Value

			invalid := []string{
				"",
				"Hello, World!",
				"Three point one four",
			}

			for _, item := range invalid {
				err := value.Set(item)
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the float64 that it references", func(t *testing.T) {
			ref := float64(3.14)
			value := newFloat64Value(&ref)

			assert.Equal(t, float64(3.14), ref)

			value.Set("3.14159")
			assert.Equal(t, float64(3.14159), ref)

			value.Set("10")
			assert.Equal(t, float64(10), ref)
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"3":          "3",
			"3.14":       "3.14",
			"3.14159":    "3.14159",
			"314.159e-2": "3.14159",
		}

		for in, expected := range inOut {
			value := new(float64Value)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}

func TestIntValue(t *testing.T) {
	t.Run("newIntValue()", func(t *testing.T) {
		intRef := 3
		intValue := newIntValue(&intRef)

		assert.Equal(t, "3", intValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value float64Value

			valid := []string{
				"3",
				"10",
				"25",
				"100000",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value intValue

			invalid := []string{
				"",
				"Hello, World!",
				"Three point one four",
				"92233720368547758070",
			}

			for _, item := range invalid {
				err := value.Set(item)
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the int that it references", func(t *testing.T) {
			ref := 5
			value := newIntValue(&ref)

			assert.Equal(t, 5, ref)

			value.Set("10")
			assert.Equal(t, 10, ref)

			value.Set("25")
			assert.Equal(t, 25, ref)
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"3":    "3",
			"10":   "10",
			"1000": "1000",
		}

		for in, expected := range inOut {
			value := new(intValue)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}

func TestIPValue(t *testing.T) {
	t.Run("newIPValue()", func(t *testing.T) {
		ipRef := net.ParseIP("127.0.0.1")
		ipValue := newIPValue(&ipRef)

		assert.Equal(t, "127.0.0.1", ipValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value ipValue

			valid := []string{
				"127.0.0.1",
				"192.168.0.1",
				"10.0.0.1",
				"255.255.255.0",
				"8.8.8.8",
				"8.8.4.4",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should error for invalid values", func(t *testing.T) {
			var value ipValue

			invalid := []string{
				"",
				"Not an IP adddress",
				"Hello, World!",
				"123 Fake Street",
				"127 0 0 1",
			}

			for _, item := range invalid {
				err := value.Set(item)
				assert.NotOK(t, err)
			}
		})

		t.Run("should modify the IP that it references", func(t *testing.T) {
			ref := net.ParseIP("127.0.0.1")
			value := newIPValue(&ref)

			assert.Equal(t, value.String(), ref.String())

			value.Set("192.168.0.1")
			assert.Equal(t, value.String(), ref.String())

			value.Set("10.0.0.1")
			assert.Equal(t, value.String(), ref.String())
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"127.0.0.1":   "127.0.0.1",
			"192.168.0.1": "192.168.0.1",
			"10.0.0.1":    "10.0.0.1",
		}

		for in, expected := range inOut {
			value := new(ipValue)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}

func TestStringValue(t *testing.T) {
	t.Run("newStringValue()", func(t *testing.T) {
		expected := "Hello, World!"
		actual := newStringValue(&expected)

		assert.Equal(t, expected, actual.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value stringValue

			valid := []string{
				"Hello",
				"World",
				"Hello, World!",
				"3.14",
				"http://www.google.co.uk/",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should modify the string that it references", func(t *testing.T) {
			ref := "Hello"

			value := newStringValue(&ref)
			assert.Equal(t, "Hello", ref)

			value.Set("World")
			assert.Equal(t, "World", ref)
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"Hello": "Hello",
			"World": "World",
			"hello": "hello",
			"world": "world",
		}

		for in, expected := range inOut {
			value := new(stringValue)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}

func TestUrlValue(t *testing.T) {
	t.Run("newURLValue()", func(t *testing.T) {
		expected := "https://www.google.co.uk/"

		actual, err := url.Parse(expected)
		assert.OK(t, err)

		actualValue := newURLValue(actual)
		assert.Equal(t, expected, actualValue.String())
	})

	t.Run("Set()", func(t *testing.T) {
		t.Run("should not error for valid values", func(t *testing.T) {
			var value urlValue

			valid := []string{
				"https://www.google.co.uk/",
				"ws://www.elliotdwright.com:9000/",
				"ftp://whouses.ftpanymore.com:21/",
			}

			for _, item := range valid {
				err := value.Set(item)
				assert.OK(t, err)
			}
		})

		t.Run("should modify the URL that it references", func(t *testing.T) {
			oldUrl := "https://www.google.co.uk/"
			newUrl := "https://www.elliotdwright.com/"

			ref, err := url.Parse(oldUrl)
			assert.OK(t, err)

			value := newURLValue(ref)
			assert.Equal(t, oldUrl, ref.String())

			value.Set(newUrl)
			assert.Equal(t, newUrl, ref.String())
		})
	})

	t.Run("String()", func(t *testing.T) {
		inOut := map[string]string{
			"https://www.google.co.uk/":        "https://www.google.co.uk/",
			"ws://www.elliotdwright.com:9000/": "ws://www.elliotdwright.com:9000/",
			"ftp://whouses.ftpanymore.com:21/": "ftp://whouses.ftpanymore.com:21/",
		}

		for in, expected := range inOut {
			value := new(urlValue)
			value.Set(in)

			actual := value.String()
			assert.Equal(t, expected, actual)
		}
	})
}
