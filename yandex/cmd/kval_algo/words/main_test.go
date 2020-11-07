package main

import "testing"


func TestField_FindFrom(t *testing.T) {
	type args struct {
		rows []string
		test string
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{
			name: "one",
			args: args{
				rows: []string{
					"cbc",
					"bab",
				},
				test: "abc",
			},
			want: 4,
		},
		{
			name: "two",
			args: args{
				rows: []string{
					"abc",
					"bbb",
					"cba",
				},
				test: "abc",
			},
			want: 4,
		},
		{
			name: "three",
			args: args{
				rows: []string{
					"aba",
					"cac",
				},
				test: "aba",
			},
			want: 6,
		},
		{
			name: "four",
			args: args{
				rows: []string{
					"aaa",
					"cbc",
				},
				test: "aaa",
			},
			want: 2,
		},
		{
			name: "ring",
			args: args{
				rows: []string{
					"zzzz",
					"bcdj",
					"ajee",
					"abcd",
				},
				test: "abcdj",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewField()
			for _, row := range tt.args.rows {
				f.addRow([]rune(row))
			}
			if got := f.Find([]rune(tt.args.test)); got != tt.want {
				t.Errorf("FindFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_Find(t *testing.T) {
	type fields struct {
		f [][]rune
	}
	type args struct {
		search []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				f: tt.fields.f,
			}
			if got := f.Find(tt.args.search); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
