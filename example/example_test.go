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
