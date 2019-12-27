package main

import "git.forms.io/legobankda/algorithms/util"

// 双指针法
func maxArea(height []int) int {
	left := 0
	right := len(height)-1
	max := util.Getmin(height[left], height[right]) * (len(height)-1)
	for left < right {
		if height[left] < height[right] {
			left ++
		} else if height[left] > height[right]{
			right --
		} else {
			left ++
			right --
		}
		max = util.Getmax(max, util.Getmin(height[left], height[right]) * (right - left))
	}
	return max
}