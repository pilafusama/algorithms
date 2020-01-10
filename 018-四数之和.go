package main

import "fmt"

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。

注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func main() {
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	ret := make([][]int, 0)
	if l < 4 {
		return ret
	}
	Quicksort2(nums, 0, l-1)
	for index1 := 0; index1 < l; index1++ {
		//if nums[index1] > target {
		//	break
		//}
		if index1 > 0 && nums[index1-1] == nums[index1] {
			continue
		}
		for index2 := index1 + 1; index2 < l; index2++ {
			//if nums[index1]+nums[index2] > target {
			//	break
			//}
			if index2 > index1+1 && nums[index2-1] == nums[index2] {
				continue
			}
			left, right := index2+1, l-1
			for left < right {
				sum := nums[index1] + nums[index2] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					ret = append(ret, []int{nums[index1], nums[index2], nums[left], nums[right]})
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				}
			}
		}
	}
	return ret
}

func Quicksort2(arr []int, left int, right int) {
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
		Quicksort2(arr, i, right)
	}
	if j > left {
		Quicksort2(arr, left, j)
	}
}
