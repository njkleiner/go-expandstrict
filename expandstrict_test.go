package expandstrict

import (
	"os"
	"testing"
)

func TestExpand(t *testing.T) {
	os.Setenv("FOO", "bar")
	os.Setenv("FIZZ", "buzz")

	s, err := Expand("Let's $FOO the ${FIZZ}", os.Getenv)

	if s != "Let's bar the buzz" {
		t.Error("invalid expansion result")
	}

	if err != nil {
		t.Errorf("unexpected error %#v", err)
	}
}

func TestExpandError(t *testing.T) {
	os.Setenv("FOO", "bar")
	os.Setenv("FIZZ", "")

	s, err := Expand("Let's $FOO the ${FIZZ}", os.Getenv)

	if s != "" {
		t.Errorf("expected empty string, actual = %s", s)
	}

	if err != ErrUndefinedVariable {
		t.Errorf("expected ErrUndefinedVariable, actual = %#v", err)
	}
}
