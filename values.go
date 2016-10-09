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

// FlagValue representats an option that does not have a value. Allows toggling behaviour.
type FlagValue interface {
	ParameterValue
	FlagValue() string
}

// boolValue abstracts functionality for parsing input that should be represented as a boolean. The
// BoolValue type also implements the FlagValue interface so that an alternative to the default
// value can be used if no value is present.
type boolValue bool

// BoolValue creates a new boolValue.
func BoolValue(ref *bool) *boolValue {
	return (*boolValue)(ref)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) String() string {
	return fmt.Sprintf("%v", *b)
}

func (b *boolValue) FlagValue() string {
	return "true"
}

// durationValue abstracts functionality for parsing input that should be represented as a
// time.Duration.
type durationValue time.Duration

// DurationValue creates a new durationValue.
func DurationValue(ref *time.Duration) *durationValue {
	return (*durationValue)(ref)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) String() string {
	return (*time.Duration)(d).String()
}

// float64Value abstracts functionality for parsing input that should be represented as a float64.
type float64Value float64

// Float64Value creates a new float64Value.
func Float64Value(ref *float64) *float64Value {
	return (*float64Value)(ref)
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

func (f *float64Value) String() string {
	return fmt.Sprintf("%v", *f)
}

// intValue abstracts functionality for parsing input that should be represented as an int.
type intValue int

// IntValue creates a new intValue.
func IntValue(ref *int) *intValue {
	return (*intValue)(ref)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) String() string {
	return fmt.Sprintf("%v", *i)
}

// ipValue abstracts functionality for parsing input that should be represented as an IP address.
type ipValue net.IP

// IPValue creates a new ipValue.
func IPValue(ref *net.IP) *ipValue {
	return (*ipValue)(ref)
}

func (s *ipValue) Set(val string) error {
	ip := net.ParseIP(val)
	if ip == nil {
		return fmt.Errorf("Invalid IP address format '%s'", val)
	}

	*s = ipValue(ip)

	return nil
}

func (s *ipValue) String() string {
	return fmt.Sprintf("%s", *s)
}

// stringValue accepts string input, and transparently assigns it to a pointer.
type stringValue string

// StringValue creates a new stringValue.
func StringValue(ref *string) *stringValue {
	return (*stringValue)(ref)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) String() string {
	return fmt.Sprintf("%s", *s)
}

type urlValue url.URL

func UrlValue(ref *url.URL) *urlValue {
	return (*urlValue)(ref)
}

func (u *urlValue) Set(val string) error {
	res, err := url.Parse(val)
	*u = urlValue(*res)
	return err
}

func (u *urlValue) String() string {
	return fmt.Sprintf("%s", *u)
}
