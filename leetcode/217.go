package leetcode

import (
	"slices"
	"sort"
)

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func containsDuplicate(nums []int) bool {
	sort.Sort(SortBy(nums))
	for {
		if len(nums) == 1 {
			return false
		}
		if _, b := slices.BinarySearch(nums[1:], nums[0]); b {
			return true
		} else {
			nums = nums[1:]
		}
	}
}

func containsDuplicateHashSet(nums []int) bool {
	hashSet := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := hashSet[nums[i]]
		if !ok {
			hashSet[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}
