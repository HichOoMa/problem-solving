package leetcode

import (
	"slices"
)

func findMaxK(nums []int) int {
	minIndex := 0
	lenght := len(nums)
	slices.Sort(nums)
	if nums[0] > 0 || nums[lenght-1] < 0 {
		return -1
	}
	// Max:
	for i := lenght - 1; nums[i] > 0 && i >= 0; i-- {
		for j := minIndex; nums[j] < 0 && j < lenght; j++ {
			if nums[i] == nums[j]*-1 {
				return nums[i]
			}
		}
	}
	return -1
}
