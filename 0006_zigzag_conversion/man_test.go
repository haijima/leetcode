package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{"PAYPALISHIRING", 3},
			want: "PAHNAPLSIIGYIR",
		}, {
			name: "Example 2",
			args: args{"PAYPALISHIRING", 4},
			want: "PINALSIGYAHRPI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
