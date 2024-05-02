package leetcode

import "testing"

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{-1, 2, -3, 3}}, 3},
		{"test 2", args{[]int{-1, 10, 6, 7, -7, 1}}, 7},
		{"test 3", args{[]int{-10, 8, 6, 7, -2, -3}}, -1},
		{"test 4", args{[]int{14, 33, 40, -33, 8, -24, -42, 30, -18, -34}}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxK(tt.args.nums); got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}
