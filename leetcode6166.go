package myleetcode

func largestPalindromic(num string) string {
	temp := make([]int, 10)
	for i := range num {
		temp[int(num[i]-'0')]++
	}

	resultArray := make([]int, len(num))
	for i := range resultArray {
		resultArray[i] = -1
	}
	start := 0
	end := len(resultArray) - 1

	for i := 9; i >= 0; i-- {
		z1 := temp[i] / 2
		s1 := start + z1
		e1 := end - z1
		for ; start < s1; start++ {
			resultArray[start] = i
		}
		for ; end > e1; end-- {
			resultArray[end] = i
		}
		temp[i] = temp[i] - 2*z1
	}

	middle := -1
	for i := 9; i >= 0; i-- {
		if temp[i] != 0 {
			middle = i
			break
		}
	}

	resultTemp := []byte{}
	for i := range resultArray {
		if resultArray[i] != -1 {
			resultTemp = append(resultTemp, byte(int('0')+resultArray[i]))
		}
	}
	for i := range resultTemp {
		//    fmt.Println(string(resultTemp[i]))
		if resultTemp[i] != '0' {
			resultTemp = resultTemp[i : len(resultTemp)-i]
			break
		}
		if i == len(resultTemp)/2-1 {
			resultTemp = []byte{}
			break
		}
	}

	if middle != -1 {
		if len(resultTemp) == 0 {
			return string([]byte{byte(middle + int('0'))})
		}
		nr := []byte{}
		middlePos := len(resultTemp) / 2
		for i := range resultTemp {
			if i == middlePos {
				nr = append(nr, byte(int('0')+middle))
			}
			nr = append(nr, resultTemp[i])
		}
		return string(nr)
	} else {
		if len(resultTemp) == 0 {
			return "0"
		}
		return string(resultTemp)
	}

}
