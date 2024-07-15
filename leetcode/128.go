package leetcode

import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	} else {
		slices.Sort(nums)
		mc := 1
		c := 1

		for i, n := range nums[1:] {
			if n == nums[i]+1 {
				c++
			} else if n == nums[i] {
				continue
			} else {
				if c > mc {
					mc = c
				}
				c = 1
			}
		}
		if c > mc {
			mc = c
		}
		return mc
	}

}
