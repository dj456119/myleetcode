/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-15 00:23:34
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-18 02:23:32
 */
package myleetcode

func removeInvalidParentheses(s string) []string {
	stc := []byte{}
	for i := range []byte(s) {

		if s[i] == ')' {
			if len(stc) != 0 && stc[len(stc)-1] == '(' {
				stc = stc[:len(stc)-1]
				continue
			}
			stc = append(stc, s[i])
		} else if s[i] == '(' {
			stc = append(stc, s[i])
		}
	}

	if len(stc) == 0 {
		return []string{s}
	}
	leftCount := 0
	rightCount := 0
	for i := range stc {
		if stc[i] == '(' {
			leftCount++
		} else {
			rightCount++
		}
	}
	return backtrace([]byte(s), 0, leftCount, rightCount, []byte{}, make(map[string]bool))
}

func backtrace(s []byte, pos int, leftCount int, rightCount int, result []byte, rm map[string]bool) []string {
	if pos == len(s) {
		if leftCount != 0 || rightCount != 0 {
			return nil
		}
		if _, ok := rm[string(result)]; ok {
			return nil
		}
		if hefa(result) {
			rm[string(result)] = true
			return []string{string(result)}
		}
		return nil
	}
	if s[pos] == '(' {
		r := []string{}
		if leftCount != 0 {
			result1 := make([]byte, len(result))
			copy(result1, result)
			r1 := backtrace(s, pos+1, leftCount-1, rightCount, result1, rm)
			if r1 != nil {
				r = append(r, r1...)
			}
		}
		result = append(result, '(')
		r2 := backtrace(s, pos+1, leftCount, rightCount, result, rm)
		if r2 != nil {
			r = append(r, r2...)
		}
		return r
	} else if s[pos] == ')' {
		r := []string{}
		if rightCount != 0 {
			result1 := make([]byte, len(result))
			copy(result1, result)
			r1 := backtrace(s, pos+1, leftCount, rightCount-1, result1, rm)
			if r1 != nil {
				r = append(r, r1...)
			}

		}
		result = append(result, ')')
		r2 := backtrace(s, pos+1, leftCount, rightCount, result, rm)
		if r2 != nil {
			r = append(r, r2...)
		}
		return r
	} else {
		result = append(result, s[pos])
		return backtrace(s, pos+1, leftCount, rightCount, result, rm)
	}
}

func hefa(s []byte) bool {
	stc := []byte{}
	for i := range s {
		if s[i] == ')' {
			if len(stc) != 0 && stc[len(stc)-1] == '(' {
				stc = stc[:len(stc)-1]
				continue
			}
			stc = append(stc, s[i])
		} else if s[i] == '(' {
			stc = append(stc, s[i])
		}
	}
	if len(stc) == 0 {
		return true
	}
	return false
}
