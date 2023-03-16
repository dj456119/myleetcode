package main

import "sort"

func maxScore(nums []int) int {
    sort.Slice(nums, func (x, y int) bool {
        return nums[x] > nums[y] 
    }) 
    sum := 0
    result := 0
    for i := range nums {
        if sum + nums[i] <= 0 {
            break
        }
        sum = sum + nums[i]
        result++
    }
    return result
}