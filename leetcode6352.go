package myleetcode

import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
	q := make([][]int, len(nums))
	for i := range nums {
		g := []int{}
		for j := range nums {
			if nums[j] != nums[i]+k && nums[j] != nums[i]-k {
				g = append(g, nums[j])
			}
		}
		q = append(q, g)
	}
	qq := make(map[string]bool)
	for i := range q {
		f := subsetsWithDup(q[i])
		for j := range f {
			sign := []byte{}
			for z := range f[j] {
				q := fmt.Sprintf("%d", f[j][z])
				sign = append(sign, []byte(q)...)
				sign = append(sign, '-')
			}
			signString := string(sign)
			qq[signString] = true
		}
	}
	return len(qq)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return backtrace6352(nums, 0, []int{})
}

func backtrace6352(nums []int, pos int, result []int) [][]int {
	if pos == len(nums) {
		for i := range result {
			result[i] = nums[result[i]]
		}
		return [][]int{result}
	}
	rr := [][]int{}
	if pos != 0 && ((len(result) != 0 && nums[pos-1] == nums[pos] && result[len(result)-1] != pos-1) || (len(result) == 0 && nums[pos-1] == nums[pos])) {
		//return nil
	} else {
		result1 := make([]int, len(result))
		copy(result1, result)
		result1 = append(result1, pos)
		r := backtrace6352(nums, pos+1, result1)
		if r != nil {
			rr = append(rr, r...)
		}
	}
	r := backtrace6352(nums, pos+1, result)
	if r != nil {
		rr = append(rr, r...)
	}

	return rr
}
