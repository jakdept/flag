package flagTrap

import (
	"fmt"
	"strings"
)

// StringFlag is a flag.Value that allows me to tell if the flag was set
// and yet I can still fall back to the default value
//
// Default should be set to the default string value of the flag
// Validator should be set to a function to validate the flag.
//
// Type is func(string)bool and it should return true on valid values
type StringTrap struct {
	string
	isSet     bool
	Default   string
	Validator func(string) bool
}

// IsSet returns whether or not if the flag was passed on the command line
func (f *StringTrap) IsSet() bool {
	return f.isSet
}

// String returns the value of the flag - either default, or what was passed
func (f *StringTrap) String() string {
	if !f.isSet || len(strings.TrimSpace(f.string)) == 0 {
		return f.Default
	}
	return f.string
}

// Set sets a given value to a flag - restricted by a given restruction function
func (f *StringTrap) Set(s string) error {
	if f.IsSet() {
		return fmt.Errorf("flag already set with value [%s]", f.string)
	}
	if len(strings.TrimSpace(s)) > 0 {
		if f.Validator != nil {
			if !f.Validator(s) {
				return fmt.Errorf("value failed validation")
			}
		}
		f.string = s
	}
	f.isSet = true
	return nil
}
