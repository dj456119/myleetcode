package myleetcode

func reverseWords(s string) string {
	buff := []byte(s)
	newBuff := []byte{}
	for i := range buff {
		if i == 0 && buff[i] == ' ' {
			continue
		}
		if i != 0 && buff[i-1] == ' ' && buff[i] == ' ' {
			continue
		}
		newBuff = append(newBuff, buff[i])
	}
	if newBuff[len(newBuff)-1] == ' ' {
		newBuff = newBuff[:len(newBuff)-1]
	}
	reve(newBuff, 0, len(newBuff)-1)
	start := 0
	for i := range newBuff {
		if i == len(newBuff)-1 {
			reve(newBuff, start, i)
			break
		}
		if newBuff[i] == ' ' {
			reve(newBuff, start, i-1)
			start = i + 1
		}
	}
	return string(newBuff)
}

func reve(buff []byte, i, j int) {
	for i < j {
		buff[i], buff[j] = buff[j], buff[i]
		i++
		j--
	}
}
