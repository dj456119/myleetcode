/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 17:47:32
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 00:03:19
 */
package myleetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	k := 0
	flag := true
	if length%2 == 0 {
		k = length / 2
	} else {
		k = length/2 + 1
		flag = false
	}
	tempz := nums1
	if len(nums1) == 0 {
		tempz = nums2
	} else if len(nums2) == 0 {
		tempz = nums1
	}
	if len(nums1) == 0 || len(nums2) == 0 {
		if flag {
			return (float64(tempz[k-1]) + float64(tempz[k])) / 2
		}
		return float64(tempz[k-1])
	}

	for k != 1 {
		pos := k/2 - 1
		a1 := 0
		b1 := 0
		if pos > len(nums1)-1 {
			a1 = len(nums1) - 1
		}
		if pos > len(nums2)-1 {
			b1 = len(nums2) - 1
		}
		if nums2[b1] > nums1[a1] {
			k = k - a1 - 1
			nums1 = nums1[a1+1:]
		} else {
			k = k - b1 - 1
			nums2 = nums2[b1+1:]
		}
		if len(nums1) == 0 || len(nums2) == 0 {
			temp := nums1
			if len(nums1) == 0 {
				temp = nums2
			} else if len(nums2) == 0 {
				temp = nums1
			}
			if flag {
				return (float64(temp[k-1]) + float64(temp[k])) / 2
			} else {
				return float64(temp[k-1])
			}
		}
	}

	if flag {
		sum := 0
		a1 := 0
		b1 := 0
		if nums1[a1] < nums2[b1] {
			sum = sum + nums1[a1]
			a1++
		} else {
			sum = sum + nums2[b1]
			b1++
		}
		if a1 > len(nums1)-1 {
			sum = sum + nums2[b1]
		}
		if b1 > len(nums2)-1 {
			sum = sum + nums1[a1]
		}
		if a1 <= len(nums1)-1 && b1 <= len(nums2)-1 {
			if nums1[a1] < nums2[b1] {
				sum = sum + nums1[a1]
			} else {
				sum = sum + nums2[b1]
			}
		}

		return float64(sum) / 2
	} else {
		return float64(min(nums1[0], nums2[0]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
