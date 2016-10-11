package console

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"
)

// ParameterValue contain the values stored in a Parameter. This interface has an implementation
// for each type of value that may be accepted.
//
// Set is used to assign the values upon parsing and validation. All input comes in as a string.
type ParameterValue interface {
	Set(string) error
	String() string
}

// FlagValue represents an option that does not have a value. Allows toggling behaviour.
type FlagValue interface {
	ParameterValue
	FlagValue() string
}

// boolValue abstracts functionality for parsing input that should be represented as a boolean. The
// BoolValue type also implements the FlagValue interface so that an alternative to the default
// value can be used if no value is present.
type boolValue bool

// BoolValue creates a new boolValue.
func newBoolValue(ref *bool) *boolValue {
	return (*boolValue)(ref)
}

// Set assigns a value to the value that this boolValue references.
func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

// String converts this boolValue to a string.
func (b *boolValue) String() string {
	return fmt.Sprintf("%v", *b)
}

// FlagValue returns the default value boolValue when no value is present (i.e. when used as a flag)
func (b *boolValue) FlagValue() string {
	return "true"
}

// durationValue abstracts functionality for parsing input that should be represented as a
// time.Duration.
type durationValue time.Duration

// DurationValue creates a new durationValue.
func newDurationValue(ref *time.Duration) *durationValue {
	return (*durationValue)(ref)
}

// Set assigns a value to the value that this durationValue references.
func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

// String converts this durationValue to a string.
func (d *durationValue) String() string {
	return (*time.Duration)(d).String()
}

// float32Value abstracts functionality for parsing input that should be represented as a float32.
type float32Value float32

// Float32Value creates a new float32Value.
func newFloat32Value(ref *float32) *float32Value {
	return (*float32Value)(ref)
}

// Set assigns a value to the value that this float32Value references.
func (f *float32Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}

	*f = float32Value(float32(v))
	return err
}

// String converts this float32Value to a string.
func (f *float32Value) String() string {
	return fmt.Sprintf("%v", *f)
}

// float64Value abstracts functionality for parsing input that should be represented as a float64.
type float64Value float64

// Float64Value creates a new float64Value.
func newFloat64Value(ref *float64) *float64Value {
	return (*float64Value)(ref)
}

// Set assigns a value to the value that this float64Value references.
func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

// String converts this float64Value to a string.
func (f *float64Value) String() string {
	return fmt.Sprintf("%v", *f)
}

// intValue abstracts functionality for parsing input that should be represented as an int.
type intValue int

// IntValue creates a new intValue.
func newIntValue(ref *int) *intValue {
	return (*intValue)(ref)
}

// Set assigns a value to the value that this intValue references.
func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

// String converts this intValue to a string.
func (i *intValue) String() string {
	return fmt.Sprintf("%v", *i)
}

// ipValue abstracts functionality for parsing input that should be represented as an IP address.
type ipValue net.IP

// IPValue creates a new ipValue.
func newIPValue(ref *net.IP) *ipValue {
	return (*ipValue)(ref)
}

// Set assigns a value to the value that this ipValue references.
func (s *ipValue) Set(val string) error {
	ip := net.ParseIP(val)
	if ip == nil {
		return fmt.Errorf("Invalid IP address format '%s'", val)
	}

	*s = ipValue(ip)

	return nil
}

// String converts this ipValue to a string.
func (s *ipValue) String() string {
	ip := net.IP(*s)

	return ip.String()
}

// stringValue accepts string input, and transparently assigns it to a pointer.
type stringValue string

// StringValue creates a new stringValue.
func newStringValue(ref *string) *stringValue {
	return (*stringValue)(ref)
}

// Set assigns a value to the value that this stringValue references.
func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

// String converts this stringValue to a string.
func (s *stringValue) String() string {
	return fmt.Sprintf("%s", *s)
}

// urlValue abstracts functionality for parsing input that should be represented as a URL.
type urlValue url.URL

// URLValue creates a new urlValue.
func newURLValue(ref *url.URL) *urlValue {
	return (*urlValue)(ref)
}

// Set assigns a value to the value that this urlValue references.
func (u *urlValue) Set(val string) error {
	res, err := url.Parse(val)
	*u = urlValue(*res)
	return err
}

// String converts this urlValue to a string.
func (u *urlValue) String() string {
	url := url.URL(*u)

	return fmt.Sprintf("%v", url.String())
}
