package myleetcode

func minimumAverageDifference(nums []int) int {
	sum := 0
	sumLeft := 0
	for i := range nums {
		sum = sum + nums[i]
	}
	sumRight := sum
	leftCount := 0
	rigthCount := len(nums)
	minSum := 99999999999
	minPos := -1
	for i := range nums {
		leftCount++
		rigthCount--
		sumLeft = sumLeft + nums[i]
		sumRight = sumRight - nums[i]
		z1 := sumLeft / leftCount
		z2 := 0
		if rigthCount == 0 {
			z2 = 0
		} else {
			z2 = sumRight / rigthCount
		}
		c := z1 - z2
		if c < 0 {
			c = -c
		}
		if c < minSum {
			minSum = c
			minPos = i
		}

	}
	return minPos
}
