/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-22 00:46:00
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-22 00:46:17
 */
/*
 * @Descripttion:请你设计一个数据结构，它能求出给定子数组内一个给定值的 频率 。

子数组中一个值的 频率 指的是这个子数组中这个值的出现次数。

请你实现 RangeFreqQuery 类：

RangeFreqQuery(int[] arr) 用下标从 0 开始的整数数组 arr 构造一个类的实例。
int query(int left, int right, int value) 返回子数组 arr[left...right] 中 value 的 频率 。
一个 子数组 指的是数组中一段连续的元素。arr[left...right] 指的是 nums 中包含下标 left 和 right 在内 的中间一段连续元素。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-22 00:46:00
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-22 00:46:01
*/
package myleetcode

type RangeFreqQuery struct {
	Arr      []int
	FirstPos map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{Arr: arr, FirstPos: make(map[int][]int)}
	for i, j := range arr {
		if _, ok := r.FirstPos[j]; ok {
			r.FirstPos[j] = append(r.FirstPos[j], i)
		} else {
			r.FirstPos[j] = []int{}
			r.FirstPos[j] = append(r.FirstPos[j], i)
		}
	}
	return r
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	//count := 0
	if pos, ok := this.FirstPos[value]; !ok {
		return 0
	} else {
		if len(pos) == 0 {
			return 0
		}
		if len(pos) >= 2 {
			if pos[len(pos)-1] < left {
				return 0
			}
			if pos[0] > right {
				return 0
			}
			if pos[0] >= left && pos[0] <= right && pos[len(pos)-1] >= left && pos[len(pos)-1] <= right {
				return len(pos)
			}
		}
		newLeft := searchInsert(pos, left)
		newRight := searchInsert(pos, right)
		if newRight >= len(pos) {
			newRight = len(pos)
		}
		if newRight < len(pos) && pos[newRight] == right {
			newRight++
		}
		if newRight > len(pos) {
			newRight = len(pos)
		}
		if newRight < newLeft {
			return 0
		} else {
			return len(pos[newLeft:newRight])
		}
		//return count
	}

}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else if nums[0] > target {
			return 0
		} else {
			return 1
		}
	}
	start := 0
	end := len(nums) - 1
	for {
		middle := (end + start) / 2
		if nums[middle] > target {
			end = middle
		} else if nums[middle] < target {
			start = middle
		} else {
			return middle
		}

		if end == start+1 {
			if nums[end] == target {
				return end
			} else if nums[end] < target {
				return end + 1
			}
		}

		if end == start || end == start+1 {
			if nums[start] == target {
				return start
			} else if nums[start] > target {
				return start
			} else {
				return start + 1
			}
		}
	}

}
