package main

import "fmt"

/*
示例 1:
输入: [2,3,1,1,4,2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,1,4,3,2,1,1,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

func main() {
	fmt.Println(canJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}

func canJump(nums []int) bool {
	if nums[0] >= len(nums)-1 {
		return true
	}
	return jump(0, nums)
}

func jump(start int, nums []int) bool {
	if nums[start] == 0 {
		if start+1 == len(nums) {
			return true
		} else {
			return false
		}
	}
	var l int
	var index int
	for i := start + 1; i <= start+nums[start]; i++ {
		if i+nums[i] > l {
			l = i + nums[i]
			index = i
			if l >= len(nums)-1 {
				return true
			}
		}
	}
	return jump(index, nums)
}
