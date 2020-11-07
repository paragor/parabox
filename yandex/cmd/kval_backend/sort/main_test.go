package main

import (
	"reflect"
	"testing"
)

func TestFilterAndSort(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "one", args: args{
			ints: []int{
				8,
				6,
				-2,
				2,
				4,
				17,
				256,
				1024,
				-17,
				-19,
			},
		}, want: []int{
			-19,
			-17,
			-2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterAndSort(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterAndSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
