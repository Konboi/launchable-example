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

func TestExampleSub(t *testing.T) {
	e := &example.Example{}

	tests := map[string]struct {
		a      int
		b      int
		expect int
	}{
		"1": {
			a:      3,
			b:      3,
			expect: 0,
		},
		"2": {
			a:      1,
			b:      4,
			expect: 3, // 1 - 4 = -3 is correct
		}}

	for c, tt := range tests {
		t.Run(c, func(t *testing.T) {
			if e.Sub(tt.a, tt.b) != tt.expect {
				t.Errorf("got:%d expect:%d", e.Sub(tt.a, tt.b), tt.expect)
			}
		})
	}
}
