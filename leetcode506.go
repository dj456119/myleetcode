/*
 * @Descripttion:给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。

运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。运动员的名次决定了他们的获奖情况：

名次第 1 的运动员获金牌 "Gold Medal" 。
名次第 2 的运动员获银牌 "Silver Medal" 。
名次第 3 的运动员获铜牌 "Bronze Medal" 。
从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。

 * @version:
 * @Author: cm.d
 * @Date: 2021-12-02 00:20:17
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-02 00:20:17
*/

package myleetcode

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))
	tempA := make(map[int]int)
	for i, j := range score {
		tempA[j] = i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	for i, j := range score {
		ti := tempA[j]
		if i == 0 {
			result[ti] = "Gold Medal"
		} else if i == 1 {
			result[ti] = "Silver Medal"
		} else if i == 2 {
			result[i] = "Bronze Medal"
		} else {
			result[i] = fmt.Sprintf("%d", i+1)
		}

	}
	return result
}

func FindRelativeRanks(score []int) []string {
	return findRelativeRanks(score)
}
