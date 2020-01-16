package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。
问总共有多少条不同的路径？
*/

func main() {
	fmt.Println(uniquePaths2(7, 3))
}

func uniquePaths(m int, n int) int {
	arr := make([][]int, 0)
	for i := 0; i < m; i++ {
		a := make([]int, n)
		arr = append(arr, a)
	}
	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}
	for i := 0; i < m; i++ {
		arr[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
	}
	return arr[m-1][n-1]
}

// 优化空间复杂度
func uniquePaths2(m int, n int) int {
	arr := make([]int, n)
	for k, _ := range arr {
		arr[k] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[j] = arr[j] + arr[j-1]
		}
	}
	return arr[n-1]
}
