package main

import "fmt"

func main() {
	fmt.Println(reverse(-123))
}

// 2147483647,-2147483648
func reverse(x int) int {
	var z int
	var y int
	for x != 0{
		y = x%10
		if z > 214748364 || (z == 214748364 && y > 7){
			return 0
		}
		if z < -214748364 || (z == -214748364 && y < -7) {
			return 0
		}
		z = z*10 + y
		x = x/10
	}
	return z
}
