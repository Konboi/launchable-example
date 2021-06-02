package example_test

import (
	"testing"

	"github.com/Konboi/launchable-example/example"
)

func TestExampleGreeting(t *testing.T) {
	e := &example.Example{}

	if e.Greeting() != "Hello" {
		t.Error("error invalid greeting")
	}
}

func TestEcho(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			"hello",
		},
		{
			"hi",
		},
	}

	e := &example.Example{}
	for _, tc := range tests {
		if e.Echo(tc.input) == tc.input {
			t.Error("error echo result invalid", "got:", e.Echo(tc.input), "expect:", tc.input)
		}
	}
}
