/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 14:40:42
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 14:51:35
 */
package myleetcode

func isValid(s string) bool {
	stack := []byte{}
	for _, j := range []byte(s) {
		if len(stack) != 0 {
			switch j {
			case '(':
				stack = push(stack, j)
			case '{':
				stack = push(stack, j)
			case '[':
				stack = push(stack, j)
			case ')':
				if stack[len(stack)-1] == '(' {
					stack, _ = pop(stack)
				} else {
					stack = push(stack, j)
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					stack, _ = pop(stack)
				} else {
					stack = push(stack, j)
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					stack, _ = pop(stack)
				} else {
					stack = push(stack, j)
				}
			}

		} else {
			if j == ']' || j == '}' || j == ')' {
				return false
			}
			stack = push(stack, j)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func pop(stack []byte) ([]byte, byte) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func push(stack []byte, a byte) []byte {
	return append(stack, a)
}
