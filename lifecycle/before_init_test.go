package lifecycle

import (
	"testing"
)

func Test_closeOld(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseOld()
		})
	}
}
