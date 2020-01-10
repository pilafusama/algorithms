package main

import "fmt"

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}

// 双指针
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	quickSort1(nums, 0, l-1)
	ret := nums[0] + nums[1] + nums[2]
	for index := 0; index < l; index++ {
		left := index + 1
		right := l - 1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if abs(ret-target) > abs(sum-target) {
				ret = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}
		}
	}
	return ret
}

func quickSort1(arr []int, left, right int) {
	start, end, tmp := left, right, arr[(left+right)/2]
	for start < end {
		for arr[start] < tmp {
			start++
		}
		for arr[end] > tmp {
			end--
		}
		if start <= end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}
	if start < right {
		quickSort1(arr, start, right)
	}
	if end > left {
		quickSort1(arr, left, end)
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
