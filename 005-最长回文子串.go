package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/

func main() {
	s := "abcda"
	fmt.Println(longestPalindrome6(s))
}

// 暴力破解  三层循环
func longestPalindrome(s string) string {
	var max int
	var test string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && j-i > max {
				test = s[i:j]
				max = j - i
			}
		}
	}
	return test
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 最长公共子串，把原来的字符串倒置了，然后找最长的公共子串就可以了。
// 求最长公共子串（不是公共子序列），有很多方法，这里用动态规划的方法。
func longestPalindrome2(s string) string {
	if s == "" {
		return s
	}
	lens := len(s)
	s2 := reverseString(s)
	arr2 := make([][]int, 0)
	for i := 0; i < lens; i++ {
		arr1 := make([]int, lens)
		arr2 = append(arr2, arr1)
	}
	var maxLen int
	var maxEnd int
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if s[i] == s2[j] {
				if i == 0 || j == 0 {
					arr2[i][j] = 1
				} else {
					arr2[i][j] = arr2[i-1][j-1] + 1
				}
				beforeRev := lens - j - 1 // 翻转前j的位置
				if arr2[i][j] > maxLen && i-beforeRev == arr2[i][j]-1 {
					maxLen = arr2[i][j]
					maxEnd = i
				}
			}
		}
	}
	return s[maxEnd-maxLen+1 : maxEnd+1]
}

//优化空间复杂度，二维数组退化为一维数组
func longestPalindrome3(s string) string {
	if s == "" {
		return s
	}
	lens := len(s)
	s2 := reverseString(s)
	arr := make([]int, lens)
	var maxLen int
	var maxEnd int
	for i := 0; i < lens; i++ {
		// 注意此处
		for j := lens - 1; j >= 0; j-- {
			if s[i] == s2[j] {
				if i == 0 || j == 0 {
					arr[j] = 1
				} else {
					arr[j] = arr[j-1] + 1
				}
			} else {
				//之前二维数组，每次用的是不同的列，所以不用置 0 。
				arr[j] = 0
			}
			beforeRev := lens - j - 1 // 翻转前j的位置
			if arr[j] > maxLen && i-beforeRev == arr[j]-1 {
				maxLen = arr[j]
				maxEnd = i
			}
		}
	}
	return s[maxEnd-maxLen+1 : maxEnd+1]
}

func reverseString(s string) string {
	arr := []rune(s)
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		arr[from], arr[to] = arr[to], arr[from]
	}
	return string(arr)
}

// 暴力破解法优化 时间复杂度：两层循环 O(n²），空间复杂度：用二维数组保存每个子串的情况 O(n²)。
func longestPalindrome4(s string) string {
	lens := len(s)
	arr2 := make([][]int, 0)
	var maxLen int
	var test string
	for i := 0; i < lens; i++ {
		arr1 := make([]int, lens)
		arr2 = append(arr2, arr1)
	}
	for l := 1; l <= lens; l++ {
		for start := 0; start < lens; start++ {
			end := start + l - 1
			if end >= lens {
				break
			}
			if (l == 1 || l == 2 || arr2[start+1][end-1] == 1) && s[start] == s[end] {
				arr2[start][end] = 1
				if l > maxLen {
					maxLen = l
					test = s[start : end+1]
				}
			}
		}
	}
	return test
}

// 暴力破解法优化 倒序
func longestPalindrome5(s string) string {
	lens := len(s)
	arr2 := make([][]int, 0)
	var maxLen int
	var test string
	for i := 0; i < lens; i++ {
		arr1 := make([]int, lens)
		arr2 = append(arr2, arr1)
	}
	for start := lens - 1; start >= 0; start-- {
		for end := start; end < lens; end++ {
			l := end - start + 1
			if (l < 3 || arr2[start+1][end-1] == 1) && s[start] == s[end] {
				arr2[start][end] = 1
				if l > maxLen {
					maxLen = l
					test = s[start : end+1]
				}
			}
		}
	}
	return test
}

// 暴力破解法优化 倒序优化，降低空间复杂度
func longestPalindrome6(s string) string {
	lens := len(s)
	arr := make([]int, lens)
	var maxLen int
	var test string
	for start := lens - 1; start >= 0; start-- {
		for end := lens - 1; end >= start; end-- {
			l := end - start + 1
			if (l < 3 || arr[end-1] == 1) && s[start] == s[end] {
				arr[end] = 1
				if l > maxLen {
					maxLen = l
					test = s[start : end+1]
				}
			} else {
				arr[end] = 0
			}
		}
	}
	return test
}
