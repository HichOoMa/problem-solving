package leetcode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"test 2", args{[]int{2, 3, 5, 0}}, []int{0, 0, 0, 30}},
		{"test 3", args{[]int{5, 9, 2, 1, 3, 9, 2, 3, 1, 7}}, []int{20412, 11340, 51030, 102060, 34020, 11340, 51030, 34020, 102060, 14580}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
