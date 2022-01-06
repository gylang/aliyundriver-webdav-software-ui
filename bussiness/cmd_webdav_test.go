package bussiness

import "testing"

func TestRunWebDav(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "111"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunWebDav()
		})
	}
}

func TestStopWebDav(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StopWebDav()
		})
	}
}
