package main

import (
	"fmt"
	"git.forms.io/legobankda/algorithms/util"
)

func main() {
	s := convert("LEETCODEISHIRING", 4)
	fmt.Println(s)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := util.Getmin(len(s), numRows)
	curRow := 0
	arr := make([][]uint8, rows)
	down := false
	for i, _ := range s {
		arr[curRow] = append(arr[curRow], s[i])
		if curRow == 0 || curRow == rows-1 {
			down = !down
		}
		if down {
			curRow++
		} else {
			curRow--
		}
	}
	ret := make([]uint8, 0)
	for i, _ := range arr {
		ret = append(ret, arr[i]...)
	}
	return string(ret)
}
