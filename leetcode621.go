/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-03 00:40:01
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-03 00:40:01
 */
package myleetcode

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	for i := range tasks {
		taskMap[tasks[i]]++
	}
	result := 0
	buffMap := make(map[byte]int)
	ntemp := n + 1
	for {
		if len(taskMap) == 0 {
			break
		}
		z1, count := findmax(taskMap, buffMap, result, n)
		if count != -1 {
			buffMap[z1] = result
			taskMap[z1]--
			if taskMap[z1] == 0 {
				delete(taskMap, z1)
			}
			result++
			ntemp--
			if ntemp == 0 {
				ntemp = n + 1
			}
			continue
		}
		result = result + ntemp
		ntemp = n + 1
	}
	return result
}

func findmax(m map[byte]int, buffMap map[byte]int, index int, n int) (byte, int) {
	var maxByte byte
	max := -1
	for k, v := range m {
		if time, ok := buffMap[k]; !ok {
			if v > max {
				max = v
				maxByte = k
			}
		} else {
			if index-time >= n+1 && v > max {
				max = v
				maxByte = k
			}
		}
	}
	return maxByte, max
}
