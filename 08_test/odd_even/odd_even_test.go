package main

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	var cases = map[string]struct {
		in  int
		out bool
	}{
		"even": {2, true},
		"odd":  {5, false},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			result := IsEven(tt.in)
			if result != tt.out {
				t.Errorf("got %v, want %v", result, tt.out)
			}
		})
	}
}
