package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{"abcabcbb"},
			want: 3,
		}, {
			name: "Example 2",
			args: args{"bbbbb"},
			want: 1,
		}, {
			name: "Example 3",
			args: args{"pwwkew"},
			want: 3,
		}, {
			name: "Example 4",
			args: args{"tmmzuxt"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
