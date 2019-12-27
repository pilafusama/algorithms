package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(twoSum3(arr, 9))
}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

//暴力破解法 时间复杂度：O(n^2)，空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

//两遍哈希表 时间复杂度：O(n)，空间复杂度：O(n)
//哈希表支持以 近似 恒定的时间进行快速查找。
//我用“近似”来描述，是因为一旦出现冲突，查找用时可能会退化到 O(n)。
//但只要你仔细地挑选哈希函数，在哈希表中进行查找的用时应当被摊销为 O(1)。
func twoSum2(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[nums[i]] = i
	}
	for i := 0; i < l; i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok && m[complement] != i {
			return []int{i, m[complement]}
		}
	}
	return []int{0, 0}
}

// 优化版
func twoSum3(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok && m[complement] != i {
			return []int{m[complement], i}
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}
