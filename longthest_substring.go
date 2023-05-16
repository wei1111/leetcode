package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var result int
	//slide := 0
	var temp int
	var leftIdx int
	for idx := 0; idx < len(s); idx++ {
		if i, ok := m[s[idx]]; ok {
			//l := idx - i
			if temp > result {
				result = temp
			}
			//slide = i+1
			if leftIdx < i {
				leftIdx = i + 1
			}
			temp = idx - leftIdx - 1
		}

		m[s[idx]] = idx
		temp++
	}
	if temp > result {
		result = temp
	}
	return result
}
