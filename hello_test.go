package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		In   string
		Want string
	}{
		{"foo", "Hello foo"},
		{"bar", "Hello bar"},
	}

	for _, test := range tests {
		t.Run(test.In, func(t *testing.T) {
			got := Hello(test.In)
			if got != test.Want {
				t.Fatalf("Hello(%q): got %q, want %q", test.In, got, test.Want)
			}
		})
	}
}
