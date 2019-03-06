package main

import (
	"fmt"
	"math"
)

// 这个好像是csp原题
//
func main()  {
	n := 0
	x := 0
	fmt.Scan(&n)
	fmt.Scan(&x)
	cur :=0
	after := []int{}
	for i :=0 ;i<n;i++ {
		fmt.Scan(&cur)
		after = append(after,cur)
	}
	before := room(x-1,after)
	for i:= 0;i<len(before);i++ {
		fmt.Print(before[i]," ")
	}
}
// 思路：记录分配次数count，往回回收，直到出现为0的某个房间，这个房间即为i，人数置为count
// 房间号从1开始
func room(x int,after []int) []int {
	//result := make([]int,n,n)
	//fmt.Println("room")
	count := 0
	for i := x; i >= 0; i-- {
		if after[i] == 0 {
			after[i] = count
			//fmt.Println("lulala")
			return after
		}
		after[i] -= 1
		count++
	}
	// 上面去掉多走的，下面开始找最小的
	min := math.MaxInt32
	index := 0
	for in, value := range after {
		if value <= min {
			min = value
			index = in
		}
	}
	if min != 0 {
		for j := 0; j < len(after); j++ {
			after[j] -= min
		}
	}
	if index != len(after)-1 {
	for j := index + 1; j < len(after); j++ {
		after[j] -= 1
		count ++
	}
}else {

	}
	after[index] = count + (min * len(after))

	return after
}