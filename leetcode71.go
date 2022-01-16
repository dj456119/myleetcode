/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-06 22:21:44
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-06 22:44:27
 */
package myleetcode

import "fmt"

func simplifyPath(path string) string {
	filter := []byte{}
	for _, j := range []byte(path) {
		if len(filter) > 0 && j == '/' && filter[len(filter)-1] == '/' {
			continue
		}
		filter = append(filter, j)
	}
	if len(filter) > 0 && filter[len(filter)-1] != '/' {
		filter = append(filter, '/')
	}

	stack := []string{}
	ns := []byte{}
	for _, j := range filter {
		if j == '/' {

			if string(ns) == "." {
				ns = []byte{}
				continue
			}
			if string(ns) == ".." {
				if len(stack) == 0 {
					ns = []byte{}
					continue
				}
				stack = stack[:len(stack)-1]
				ns = []byte{}
				continue
			}
			if len(ns) == 0 {
				continue
			}
			stack = append(stack, string(ns))
			ns = []byte{}
		} else {
			ns = append(ns, j)

		}
	}
	result := ""
	if len(stack) == 0 {
		return "/"
	}
	for _, j := range stack {
		result = fmt.Sprintf("%s/%s", result, j)
	}
	return result
}
