package main

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0
示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
 */
func main() {
	A := []int{2}
	B := []int{1}
	fmt.Println(findMedianSortedArrays(A, B))
}

// 二分法，先考虑普遍情况，再考虑特殊情况
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		temp := nums1
		nums1 = nums2
		nums2 = temp
		m = len(nums1)
		n = len(nums2)
	}
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax--
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin++
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = getmax2(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = getmin2(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0
}

func getmax2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getmin2(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
