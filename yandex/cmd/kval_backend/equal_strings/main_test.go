package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_isEqualsRows(t *testing.T) {
	type args struct {
		a Row
		b Row
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				a: Row{
					Cells: []valueType{hashValue("one"), emptyValue},
				},
				b: Row{
					Cells: []valueType{hashValue("one"), emptyValue},
				},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				a: Row{
					Cells: []valueType{hashValue("one"), hashValue("one")},
				},
				b: Row{
					Cells: []valueType{hashValue("one"), emptyValue},
				},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				a: Row{
					Cells: []valueType{hashValue("two"), hashValue("one")},
				},
				b: Row{
					Cells: []valueType{hashValue("one"), emptyValue},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqualsRows(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isEqualsRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashPair(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one", args: args{
			a: 3,
			b: 4,
		},
			want: "3_4",
		},
		{
			name: "one", args: args{
			a: 4,
			b: 3,
		},
			want: "3_4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashPair(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hashPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testOne(t *testing.T) {
	db := NewDataBase([]int{1, 2})

	t.Run("one", func(t *testing.T) {
		for _, row := range [][]string{
			{"a", "b", ""},
			{"", "b", "c"},
			{"a", "c", ""},
			{"a", "b", "d e"},
			{"", "", "d e"},
			{"a", "", "f"},
		} {
			db.AddRow(row)
		}
		want := map[string][]int{
			"0_1": {0, 1},
			"0_3": {0, 3},
			"3_4": {3, 4},
		}
		if got := db.FindDuplicates(); !reflect.DeepEqual(got, want) {
			t.Errorf("FindDuplicates() = %v, want %v", got, want)
		}
	})
}

func Test_hashValue(t *testing.T) {
	fmt.Println(hashValue("onesdasdasdasjdkjlasjasdasdasdlkdjaskld"))
	fmt.Println(hashValue("one"))
}
