/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-05 00:04:31
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-05 00:24:07
 */
package myleetcode

func modifyString(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			result = append(result, s[i])
		} else {
			if i == 0 {
				if len(s) > 1 {
					if s[i+1] == 'a' {
						result = append(result, 'z')
					} else {
						result = append(result, 'a')
					}
				} else {
					result = append(result, 'a')
				}

			} else {

				if result[i-1] == 'z' {
					result = append(result, 'a')
				} else {
					result = append(result, result[i-1]+1)
				}

				if i+1 < len(s) && result[i] == s[i+1] {
					if result[i] == 'z' {
						result[i] = 'a'
					} else {
						result[i]++
					}
				}
			}
		}
	}
	return string(result)
}
