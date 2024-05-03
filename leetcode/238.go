package leetcode

// func productExceptSelf(nums []int) []int {
// 	lengh := len(nums)
// 	lastIndex := lengh - 1
// 	direct := make([]int, lengh)
// 	indirect := make([]int, lengh)
// 	// for index, val := range nums {
// 	// 	if index > 0 && index < lengh-1 {
// 	// 		direct[index] = direct[index-1] * val
// 	// 		indirect[lengh-index-1] = nums[lengh-index] * nums[lengh-index-1]
// 	// 	} else if index == 0 {
// 	// 		direct[0] = 1
// 	// 		indirect[lengh-1] = 1
// 	// 	} else if index == lengh-1 {
// 	// 		direct[index] = val
// 	// 		indirect[lengh-1] = nums[0]
// 	// 	}
// 	// }
// 	for i := 0; i < lengh; i++ {
// 		if i == 0 {
// 			direct[0] = nums[0]
// 			indirect[lastIndex] = nums[lastIndex]
// 		} else {
// 			direct[i] = direct[i-1] * nums[i]
// 			indirect[lastIndex-i] = indirect[lastIndex-i+1] * nums[lastIndex-i]
// 		}
// 	}
// 	answer := make([]int, lengh)
// 	for i := 0; i < lengh; i++ {
// 		if i == 0 {
// 			answer[i] = indirect[i]
// 		} else if i == lengh-1 {
// 			answer[i] = direct[i-1]
// 		} else {
// 			answer[i] = direct[i-1] * indirect[i-1]
// 		}
// 	}
// 	return answer
// }

func productExceptSelf(nums []int) []int {
	lenght := len(nums)
	answer := make([]int, lenght)
	direct := 1
	indirect := 1
	for i := 0; i < lenght; i++ {
		if i <= lenght-1-i {
			answer[i] = direct
		} else {
			answer[i] = answer[i] * direct
		}
		direct = direct * nums[i]
		if i < lenght-1-i {
			answer[lenght-1-i] = indirect
		} else {
			answer[lenght-1-i] = answer[lenght-1-i] * indirect
		}
		indirect = indirect * nums[lenght-1-i]
	}
	return answer
}
