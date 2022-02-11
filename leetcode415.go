/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 01:25:59
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-12 01:25:59
 */
package myleetcode

func addStrings(num1 string, num2 string) string {
	return string(add([]byte(num1), []byte(num2)))
}

func add(s1, s2 []byte) []byte {
	result := []byte{}
	s1, s2 = short(s1, s2)
	flag := 0
	m := len(s2) - 1
	for i := len(s1) - 1; i >= 0; i-- {
		z := flag + int(s1[i]-'0') + int(s2[m]-'0')
		flag = z / 10
		result = append(result, byte(z%10)+'0')
		m--
	}
	j := len(s2) - len(s1) - 1
	for j >= 0 {
		z := flag + int(s2[j]-'0')
		flag = z / 10
		result = append(result, byte(z%10)+'0')
		j--
	}
	if flag != 0 {
		result = append(result, byte(flag)+'0')
	}
	return reverse(result)
}

func short(s1, s2 []byte) ([]byte, []byte) {
	if len(s1) < len(s2) {
		return s1, s2
	}
	return s2, s1
}

func reverse(s []byte) []byte {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
