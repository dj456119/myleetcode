/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-02 11:27:31
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-02 11:31:13
 */
package myleetcode

func checkString(s string) bool {
	fi := 0
	bflag := false
	for i, j := range s {
		if j == 'a' {
			if bflag {
				return false
			}
			fi = i
		}
		if j == 'b' {
			bflag = true
		}
	}
	if fi == 0 {
		return true
	}

}
