package main

import "fmt"

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

func main() {
	arr := []int{1, 2, 3}
	nextPermutation2(arr)
	fmt.Println(arr)
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for i > 0 {
		if nums[i] > nums[i-1] {
			// 找到比nums[i-1]大的数据，交换后，翻转
			for j := len(nums) - 1; j >= i; j-- {
				if nums[i-1] < nums[j] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					reverse1(nums, i)
					return
				}
			}
		}
		i--
	}
	reverse1(nums, 0)
}

func nextPermutation2(nums []int) {
	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i > 0 {
		j := len(nums) - 1
		for j >= i && nums[i-1] >= nums[j] {
			j--
		}
		nums[i-1], nums[j] = nums[j], nums[i-1]
	}
	reverse1(nums, i)
}

func reverse1(arr []int, start int) {
	l := start
	r := len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
