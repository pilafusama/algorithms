package main

import (
	"fmt"
	"strings"
)

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
*/

func main() {
	a := "leeeeetcode"
	c := ""
	b := isSubsequence(a, c)
	fmt.Println(b)
}

func isSubsequence(t string, s string) bool {
	a := 0
	for i := 0; i < len(t); i++ {
		index := strings.Index(s[a:], t[i:i+1]) + 1
		if index != 0 {
			a += index
		} else {
			return false
		}
	}
	return true
}

func isSubsequence2(t string, s string) bool {
	i, j := 0, 0
	for i < len(t) && j < len(s) {
		if t[i] == s[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == len(t) {
		return true
	}
	return false
}
