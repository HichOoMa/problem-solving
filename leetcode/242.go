package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		sHash := make(map[byte]int, len(s))
		tHash := make(map[byte]int, len(s))
		for i := 0; i < len(s); i++ {
			sChar := s[i]
			if _, ok := sHash[sChar]; ok {
				sHash[sChar]++
			} else {
				sHash[sChar] = 1
			}
			tChar := t[i]
			if _, ok := tHash[tChar]; ok {
				tHash[tChar]++
			} else {
				tHash[tChar] = 1
			}
		}
		if len(sHash) != len(tHash) {
			return false
		}
		for key, value := range sHash {
			if tHash[key] != value {
				return false
			}
		}
		return true
	}
}
