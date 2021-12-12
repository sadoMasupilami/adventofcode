package main

import "testing"

func Test_containsAllLetters(t *testing.T) {
	type args struct {
		s      string
		search string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "a", args: struct {
			s      string
			search string
		}{s: "abcde", search: "abcde"}, want: true},
		{name: "b", args: struct {
			s      string
			search string
		}{s: "abcde", search: "bd"}, want: true},
		{name: "c", args: struct {
			s      string
			search string
		}{s: "abcde", search: "bcf"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAllLetters(tt.args.s, tt.args.search); got != tt.want {
				t.Errorf("containsAllLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
