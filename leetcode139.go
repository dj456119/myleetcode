/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-02 00:58:22
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-02 01:30:57
 */
package myleetcode

func wordBreak(s string, wordDict []string) bool {
	dpmap := make(map[string]bool)
	dparr := make([]bool, len(s))
	for _, j := range wordDict {
		dpmap[j] = true
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if _, ok := dpmap[s[i:j+1]]; ok && (i == 0 || dparr[i-1] == true) {
				dparr[j] = true
			}
		}
	}
	return dparr[len(dparr)-1]
}
