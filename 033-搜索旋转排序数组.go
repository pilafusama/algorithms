package main

import "fmt"

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。

示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4

示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

func main() {
	fmt.Println(search([]int{5, 1, 3}, 5))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return search2(nums, 0, len(nums)-1, target)
}

func search2(nums []int, left, right, target int) int {
	if left == right && nums[left] != target {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < nums[left] { // 左边断了
		// 右边二分查找
		if nums[mid] < target && nums[right] >= target {
			return midSearch2(nums, mid+1, right, target)
		} else {
			// 去左边找
			return search2(nums, left, mid-1, target)
		}
	} else { // 右边断了
		// 左边二分查找
		if nums[mid] > target && nums[left] <= target {
			return midSearch2(nums, left, mid-1, target)
		} else {
			// 去右边找
			return search2(nums, mid+1, right, target)
		}
	}
}

// 二分法查找
func midSearch(arr []int, left, right, target int) int {
	if left == right && arr[left] != target {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return midSearch(arr, left, mid-1, target)
	} else {
		return midSearch(arr, mid+1, right, target)
	}
}

// 二分法查找优化
func midSearch2(arr []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
