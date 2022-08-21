package myleetcode

func minimumOperations(nums []int) int {
	result := 0
	for {
		flag := false
		min := 9999999
		for i := range nums {
			if nums[i] != 0 {
				flag = true
			}
			if nums[i] < min && nums[i] != 0 {
				min = nums[i]
			}
		}
		if !flag {
			return result
		}
		result++
		//  fmt.Println(nums, min)
		for i := range nums {
			if nums[i] == 0 {
				continue
			}
			nums[i] = nums[i] - min
		}
	}
}
