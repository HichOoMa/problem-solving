package leetcode

func topKFrequent(nums []int, k int) []int {
	mapOcc := make(map[int]int)
	for _, i := range nums {
		if _, ok := mapOcc[i]; ok {
			mapOcc[i]++
		} else {
			mapOcc[i] = 1
		}
	}

	occ := []int{}
	res := []int{}
	for key, val := range mapOcc {
		if len(res) < k {
			occ = append(occ, val)
			res = append(res, key)
		} else {
			for index, _ := range res {
				if occ[index] < val {
					valaux := occ[index]
					keyaux := res[index]
					occ[index] = val
					res[index] = key
					val = valaux
					key = keyaux
				}
			}
		}
	}
	return res
}
