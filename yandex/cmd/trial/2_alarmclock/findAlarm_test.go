package main

import (
	"testing"
)

func Test_FindAlarms(t *testing.T) {
	type args struct {
		clocks        []int
		delta         int
		clockRequired int
	}
	tests := []struct {
		name     string
		args     args
		wantTime int
	}{
		{name: "first", args: args{
			clocks:        []int{1, 2, 3, 4, 5, 6},
			delta:         5,
			clockRequired: 10,
		}, wantTime: 10},
		{name: "first", args: args{
			clocks:        []int{1},
			delta:         100,
			clockRequired: 1,
		}, wantTime: 1},
		{name: "first", args: args{
			clocks:        []int{5, 22, 17, 13, 8,100,100,300},
			delta:         7,
			clockRequired: 12,
		}, wantTime: 27},
		{name: "first", args: args{
			clocks:        []int{5, 22, 17, 13, 8},
			delta:         7,
			clockRequired: 12,
		}, wantTime: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTime := FindAlarms(tt.args.clocks, tt.args.delta, tt.args.clockRequired); gotTime != tt.wantTime {
				t.Errorf("findAlarms() = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}
