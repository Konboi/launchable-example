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
func TestExampleAdd(t *testing.T) {
	e := &example.Example{}

	tests := map[string]struct {
		a      int
		b      int
		expect int
	}{
		"1": {
			a:      3,
			b:      3,
			expect: 6,
		},
		"2": {
			a:      3,
			b:      4,
			expect: 7,
		}}

	for c, tt := range tests {
		t.Run(c, func(t *testing.T) {
			if e.Add(tt.a, tt.b) != tt.expect {
				t.Errorf("add:%d expect:%d", e.Add(tt.a, tt.b), tt.expect)
			}
		})
	}
}

func TestExampleDiv(t *testing.T) {
	e := &example.Example{}

	tests := map[string]struct {
		a      int
		b      int
		expect int
	}{
		"1": {
			a:      3,
			b:      3,
			expect: 1,
		},
		"2": {
			a:      8,
			b:      4,
			expect: 1, // 2 is correct
		}}

	for c, tt := range tests {
		t.Run(c, func(t *testing.T) {
			if e.Div(tt.a, tt.b) != tt.expect {
				t.Errorf("add:%d expect:%d", e.Add(tt.a, tt.b), tt.expect)
			}
		})
	}
}
