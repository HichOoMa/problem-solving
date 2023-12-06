package main

import (
	"testing"
)

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1(tt.args.lines)
		})
	}
}

func Test_isInt(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"valid int", "1", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInt(tt.args); got != tt.want {
				t.Errorf("isInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFullNumber(t *testing.T) {
	type args struct {
		str   string
		index int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getFullNumber(tt.args.str, tt.args.index)
			if got != tt.want {
				t.Errorf("getFullNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getFullNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
