package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Ares"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Ares") = %q, %v want to match for %q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	message, err := Hello(name)

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v want to match "", error`, message, err)
	}
}
