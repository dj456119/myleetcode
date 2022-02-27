/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-26 01:07:43
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-28 01:07:54
 */
package myleetcode

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dparri_1 := 1
	dparri_0 := 1
	dparri__1 := 1
	for i := range s {
		if i == 0 {
			continue
		} else {
			dparri__1 = dparri_0
			ii := int(s[i]-'0') + int(s[i-1]-'0')*10
			if s[i] == '0' && ii > 26 {
				return 0
			}
			if i != len(s)-1 && s[i+1] == '0' {
				continue
			}
			if s[i] == '0' && s[i-1] == '0' {
				return 0
			}
			if s[i] == '0' || s[i-1] == '0' {
				continue
			}
			if ii <= 26 {
				dparri__1 = dparri__1 + dparri_1

			}
			dparri_1 = dparri_0
			dparri_0 = dparri__1
		}

	}
	return dparri__1

}
