package leetcode

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{[]int{1, 2, 3, 1}}, true},
		{"test 2", args{[]int{1, 2, 3, 4}}, false},
		{"test 3", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicateHashSet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{[]int{1, 2, 3, 1}}, true},
		{"test 2", args{[]int{1, 2, 3, 4}}, false},
		{"test 3", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
		{"test 4", args{[]int{3, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicateHashSet(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicateHashSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
