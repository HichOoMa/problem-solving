package leetcode

import (
	"sort"
)

// func groupAnagrams(strs []string) [][]string {
// 	mapCateg := make(map[int][]string)
// 	for _, str := range strs {
// 		if _, ok := mapCateg[len(str)]; ok {
// 			mapCateg[len(str)] = append(mapCateg[len(str)], str)
// 		} else {
// 			mapCateg[len(str)] = []string{str}
// 		}
// 	}
// 	mapReturn := make(map[int][]string)
// 	for _, values := range mapCateg {
// 		for str := range values {

//			}
//		}
//	}
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	mapReturn := make(map[string]int)
	returnData := [][]string{}
	for _, str := range strs {
		strRune := []rune(str)
		sort.Sort(sortRunes(strRune))
		if val, ok := mapReturn[string(strRune)]; ok {
			returnData[val] = append(returnData[val], str)
		} else {
			mapReturn[string(strRune)] = len(mapReturn)
			returnData[val] = []string{str}
		}
	}
	return returnData
}
