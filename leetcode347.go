/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-09 00:43:21
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-09 00:49:35
 */
package myleetcode

func topKFrequent(nums []int, k int) []int {
	z := make(map[int]int)
	for i := range nums {
		z[nums[i]]++
	}
	countMap := make(map[int][]int)
	for k, v := range z {
		if h, ok := countMap[v]; ok {
			h = append(h, k)
		} else {
			h = []int{k}
			countMap[v] = h
		}
	}
	h := []int{}
	for i := range nums {
		h = Push(h, nums[i])
		if len(h) > k {
			h, _ = Pop(h)
		}
	}
	result := []int{}
	for i := range h {
		t := countMap[h[i]]
		result = append(result, t...)
	}
	return result
}

func Push(h []int, num int) []int {
	h = append(h, num)
	if len(h) == 1 {
		return h
	}
	return Up(h, len(h)-1)
}

func Up(h []int, pos int) []int {
	i := pos
	for i != 0 && h[i] < h[(i-1)/2] {
		h[i], h[(i-1)/2] = h[(i-1)/2], h[i]
		i = (i - 1) / 2
	}
	return h
}

func Down(h []int, pos int) []int {
	for {
		left := 2*pos + 1
		right := 2*pos + 2
		minPos := -1
		if right > len(h)-1 {
			if left > len(h)-1 {
				return h
			}
			minPos = left
		} else {
			if h[left] < h[right] {
				minPos = left
			} else {
				minPos = right
			}
		}

		if h[minPos] < h[pos] {
			h[minPos], h[pos] = h[pos], h[minPos]
			pos = minPos
		} else {
			return h
		}
	}
}

func Remove(h []int, pos int) []int {
	if len(h) == 0 {
		return h
	}
	if pos > len(h)-1 {
		return h
	}
	if pos != len(h)-1 {
		h[pos], h[len(h)-1] = h[len(h)-1], h[pos]
	}
	h = h[:len(h)-1]
	h = Up(h, pos)
	h = Down(h, pos)
	return h
}

func Pop(h []int) ([]int, int) {
	z := h[0]
	h = Remove(h, 0)
	return h, z
}
