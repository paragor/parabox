package main

import (
	"testing"
)

func Test_isUpper(t *testing.T) {
	type args struct {
		x rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one1",
			args: args{x: 'A'},
			want: true,
		},
		{
			name: "one2",
			args: args{x: 'F'},
			want: true,
		},
		{
			name: "one3",
			args: args{x: 'Z'},
			want: true,
		},
		{
			name: "one4",
			args: args{x: 'a'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUpper(tt.args.x); got != tt.want {
				t.Errorf("isUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		s []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the0",
			args: args{
				s: []rune("APPLE II"),
			},
			want: 10,
		},
		{
			name: "one1",
			args: args{
				s: []rune("h"),
			},
			want: 1,
		},
		{
			name: "one2",
			args: args{
				s: []rune("H"),
			},
			want: 2,
		},
		{
			name: "one3",
			args: args{
				s: []rune("Hello World"),
			},
			want: 13,
		},
		{
			name: "one4",
			args: args{
				s: []rune("APPLE II"),
			},
			want: 10,
		},
		{
			name: "one4",
			args: args{
				s: []rune("APPLE II I"),
			},
			want: 13,
		},
		{
			name: "one5",
			args: args{
				s: []rune("   "),
			},
			want: 3,
		},
		{
			name: "one6",
			args: args{
				s: []rune(" "),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.s); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

