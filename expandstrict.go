// Package expandstrict provides a replacement for os.Expand that returns
// an error if any of the variables to be expanded are undefined.
package expandstrict

import (
	"errors"
	"os"
)

var (
	ErrUndefinedVariable = errors.New("cannot expand undefined variable")
)

func wrap(mapping func(string) string) func(string) string {
	return func(s string) string {
		if rep := mapping(s); rep != "" {
			return rep
		}

		return "X"
	}
}

// Expand mimics the behavior of os.Expand but returns ErrUndefinedVariable
// if any of the variables to be expanded are undefined.
func Expand(s string, mapping func(string) string) (string, error) {
	exp := os.Expand(s, mapping)
	min := os.Expand(s, wrap(mapping))

	if len(exp) < len(min) {
		return "", ErrUndefinedVariable
	}

	return exp, nil
}
