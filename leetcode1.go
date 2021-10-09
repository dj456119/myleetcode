/**
数组中两数之和等于第三个数
*/
package myleetcode

func twoSum(nums []int, target int) []int {
	dataMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := dataMap[target-num]; ok {
			return []int{v, i}
		}
		dataMap[num] = i
	}
	return nil
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
