package main

import (
	"fmt"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	lens := len(nums)
	ret := make([][]int, 0)
	if lens < 3 {
		return ret
	}
	// abcdefg acf
	Quicksort(nums, 0, lens-1)
	for i := 0; i < lens; i++ {
		// todo 减少无谓的计算
		if nums[i] > 0 {
			break
		}
		// todo i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := lens - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				// todo left去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// todo right去重
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}

// quick sort
func Quicksort(arr []int, left, right int) {
	i := left
	j := right
	temp := arr[(left+right)/2]
	for i <= j {
		for arr[i] < temp {
			i++
		}
		for arr[j] > temp {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if i < right {
		Quicksort(arr, i, right)
	}
	if j > left {
		Quicksort(arr, left, j)
	}
}
