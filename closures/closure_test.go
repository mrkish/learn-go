package main

import (
	"testing"
)

func Test_adder(t *testing.T) {

	tests := []struct {
		name string
		want func(int) int
	}{
		{
			name: "Wtf",
			want: func(int) int {
				return 0
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		//	if got := adder(); !reflect.DeepEqual(got, tt.want) {
		//		t.Errorf("adder() = %v, want %v", got, tt.want)
		//	}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "wtf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}