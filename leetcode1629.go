/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-09 00:09:45
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-10 22:30:29
 */
package myleetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {

	maxTime := 0
	index := 0
	for i := range releaseTimes {
		if i == 0 {

			maxTime = releaseTimes[0]
		} else {
			t := releaseTimes[i] - releaseTimes[i-1]
			if t > maxTime {

				maxTime = t
				index = i
			} else if t == maxTime {
				if keysPressed[i] > keysPressed[index] {
					index = i
				}
			}

		}

	}
	return keysPressed[index]
}
