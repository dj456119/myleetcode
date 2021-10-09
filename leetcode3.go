/**
无重复字符最长子串
*/
package myleetcode

func lengthOfLongestSubstring(s string) int {
	start := 0
	dataMap := make(map[rune]int)
	maxLength := 0
	flag := true
	for i, c := range s {
		if pos, ok := dataMap[c]; ok {
			flag = false
			if pos < start {

			} else if nl := i - start; nl > maxLength {
				maxLength = nl
				start = pos + 1
			} else {
				start = pos + 1
			}
			dataMap[c] = i
		} else {

			dataMap[c] = i
		}
	}
	if flag {
		return len(s)
	}
	lastLength := len(s) - start
	if lastLength > maxLength {
		return lastLength
	}
	return maxLength
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}
