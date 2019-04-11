package main

import "fmt"

// 找到N的二进制表示中两个连续的1最远的距离
// 求和2的最大公约数
func binaryGap(N int) int {
	max := -1
	cur := 0
	for N > 0 {
		a := N % 2
		N = N / 2
		if a == 0 && max != -1 {
			cur++
		} else if a == 1 {
			if cur > max {
				max = cur
			}
			cur = 1
		}
	}
	return max
}
func main() {
	fmt.Println(binaryGap(8))
}
