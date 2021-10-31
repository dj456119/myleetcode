/*
 * @Descripttion:给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-01 00:54:06
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-01 01:06:35
 */

package myleetcode

func distributeCandies(candyType []int) int {
	countMap := make(map[int]int)

	for _, j := range candyType {
		if _, ok := countMap[j]; !ok {
			countMap[j] = 1
		}
		if len(countMap) == (len(candyType) / 2) {
			return len(countMap)
		}
	}
	return len(countMap)
}

func DistributeCandies(candyType []int) int {
	return distributeCandies(candyType)
}
