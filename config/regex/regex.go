package regex

import (
	"regexp"
)

type Expression struct {
	*regexp.Regexp
}

// UnmarshalYAML implements the unmarshaler interface for yaml.
func (r *Expression) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	regex, err := regexp.Compile(s)
	if err != nil {
		return err
	}
	r.Regexp = regex
	return nil
}

// MarshalYAML implements the marshaler interface for yaml.
func (r Expression) MarshalYAML() (interface{}, error) {
	if r.Regexp != nil {
		return r.Regexp.String(), nil
	}
	return nil, nil
}
