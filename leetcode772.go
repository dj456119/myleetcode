/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-25 00:50:10
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-26 00:21:30
 */
package myleetcode

import (
	"container/list"
	"fmt"
)

func calculate772(s string) int {
	opermap := make(map[byte]int)
	opermap['+'] = 10
	opermap['-'] = 10
	opermap['*'] = 50
	opermap['/'] = 50
	opermap[')'] = 60
	opermap['#'] = 0
	s = s + "#"
	resultLsit := list.New()
	operStack := list.New()

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if 0 <= (s[i]-'0') && (s[i]-'0') <= 9 {
			num := 0
			for {
				c := s[i]
				if 0 <= (c-'0') && (c-'0') <= 9 && i < len(s) {
					num = num*10 + int(c-'0')
					i++
				} else {
					i--
					break
				}
			}
			resultLsit.PushBack(num)
			continue
		} else {
			if operStack.Len() == 0 || s[i] == '(' {
				operStack.PushBack(s[i])
				continue
			} else {
				operTemp := operStack.Back()
				oper := operTemp.Value.(byte)
				if s[i] != ')' && (opermap[oper] < opermap[byte(s[i])]) {
					operStack.PushBack(s[i])
				} else if s[i] == ')' {
					for {
						operTemp := operStack.Back()
						oper := operTemp.Value.(byte)
						if oper == '(' {
							operStack.Remove(operTemp)
							break
						} else {
							operStack.Remove(operTemp)
							resultLsit.PushBack(oper)
						}

					}
				} else {
					for {
						if operStack.Len() == 0 {
							operStack.PushBack(s[i])
							break
						}
						operTemp := operStack.Back()
						oper := operTemp.Value.(byte)
						if opermap[oper] >= opermap[byte(s[i])] {
							operStack.Remove(operTemp)
							if oper != '(' {
								resultLsit.PushBack(oper)
							}
							continue
						} else {
							operStack.PushBack(s[i])
							break
						}

					}
				}
			}
		}

	}
	return calculateRePorland(resultLsit)

}

func calculateRePorland772(rePorlandList *list.List) int {
	calStack := list.New()
	for rePorlandList.Len() != 0 {
		temp := rePorlandList.Front()
		switch temp.Value.(type) {
		case int:
			calStack.PushBack(temp.Value)
		case byte:
			temp2 := calStack.Back()
			temp2Num := temp2.Value.(int)
			calStack.Remove(temp2)
			temp1 := calStack.Back()
			temp1Num := temp1.Value.(int)
			calStack.Remove(temp1)
			result := 0
			switch temp.Value.(byte) {
			case '+':
				result = temp1Num + temp2Num
			case '-':
				result = temp1Num - temp2Num
			case '*':
				result = temp1Num * temp2Num
			case '/':
				result = temp1Num / temp2Num
			}
			calStack.PushBack(result)
		default:
			fmt.Println("不认识", temp.Value)
		}
		rePorlandList.Remove(temp)
	}
	return calStack.Front().Value.(int)
}

func Calculate772(s string) int {
	return calculate772(s)
}
