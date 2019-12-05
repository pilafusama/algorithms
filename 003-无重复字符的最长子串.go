package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func main() {
	s := "abcdefgcdefg"
	fmt.Println(lengthOfLongestSubstring2(s))
}

// 简单粗暴迭代法
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	arr := strings.Split(s, "")
	str := arr[0]
	for i := 1; i < len(arr); i++ {
		if !strings.Contains(str, arr[i]) {
			str += arr[i]
		} else {
			index := strings.Index(str, arr[i])
			len := lengthOfLongestSubstring(s[index+1:])
			if i < len {
				return len
			} else {
				return i
			}
		}
	}
	return len(str)
}

// 滑动窗口，动态的记录left和max值
func lengthOfLongestSubstring2(s string) int {
	if s == "" {
		return 0
	}
	m := make(map[rune]int)
	left := 0
	max := 0
	for index, c := range s {
		if index2, ok := m[c]; ok {
			left = getmax(left, index2+1)
		}
		m[c] = index
		max = getmax(max, index-left+1)
	}
	return max
}

func getmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
