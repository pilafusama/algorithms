package main

import (
	"fmt"
	"git.forms.io/legobankda/algorithms/util"
	"strings"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:所有输入只包含小写字母 a-z 。
*/

func main() {
	arr := []string{"aa"}
	fmt.Printf(longestCommonPrefix(arr))
}

// 暴力法
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for j := 1; j <= len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], strs[0][0:j]) {
				return strs[0][0 : j-1]
			}
		}
	}
	return strs[0]
}

// LCP(S1…Sn)=LCP(LCP(LCP(S1,S2),S3),…Sn)
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return prefix
			}
		}
	}
	return prefix
}

// 逐列比较
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

// 分治(不断拆分迭代)
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefixSub(strs, 0, len(strs)-1)
}

func longestCommonPrefixSub(strs []string, left int, right int) string {
	if left == right {
		return strs[left]
	}
	mid := (left + right) / 2
	strLeft := longestCommonPrefixSub(strs, left, mid)
	strRight := longestCommonPrefixSub(strs, mid+1, right)
	return longestCommonPrefixSub2(strLeft, strRight)
}

func longestCommonPrefixSub2(s1, s2 string) string {
	minLens := util.Getmin(len(s1), len(s2))
	for i := 0; i < minLens; i++ {
		if s1[i] != s2[i] {
			return s1[0:i]
		}
	}
	return s1[0:minLens]
}
