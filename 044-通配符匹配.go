package main

import "fmt"

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

输入:
s = "adceb"
p = "*a*b"
输出: true

输入:
s = "acdcb"
p = "a*c?b"
输入: false
*/

func main() {
	s := "adceb"
	p := "*a*b"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	arr := make([][]bool, 0)
	for i := 0; i <= m; i++ {
		a := make([]bool, n+1)
		arr = append(arr, a)
	}
	arr[0][0] = true
	for i := 1; i <= n; i++ {
		arr[0][i] = (arr[0][i-1] && (p[i-1:i] == "*"))
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1:j] == "?" {
				arr[i][j] = arr[i-1][j-1]
			}
			if p[j-1:j] == "*" {
				arr[i][j] = arr[i][j-1] || arr[i-1][j]
			}
		}
	}
	return arr[m][n]
}
