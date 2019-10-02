package closures

import (
	"reflect"
	"testing"
)

func Test_adder(t *testing.T) {
	tests := []struct {
		name string
		want func(int) int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adder() = %v, want %v", got, tt.want)
			}
		})
	}
}