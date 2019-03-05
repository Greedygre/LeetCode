package main

import (
	"fmt"
	"sort"
)

// 从一个任意大小写英文字符串中找其中最长连续的英文字符串（不区分大小写）
func main() {
	result := []int{}
	n := 0
	fmt.Scan(&n)
	//love := []int{}
	love := make(map[int][]int)

	pre := 0
	for i := 0; i < n; i++ {

		fmt.Scan(&pre)
		if _, ok := love[pre]; ok {
			//存在
			cur := love[pre]
			cur = append(cur, i+1)
			love[pre] = cur
			continue
		}
		//如果不存在
		love[pre] = []int{i + 1}

	}

	q := 0
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		start := 0
		fmt.Scan(&start)

		end := 0
		fmt.Scan(&end)
		lovepoint := 0
		fmt.Scan(&lovepoint)

		result = append(result, search(start, end, love[lovepoint]))

	}
	for _, i := range result {
		fmt.Println(i)
	}

}

func search(start, end int, love []int) int {
	sort.Ints(love)
	count := 0
	for _, value := range love {
		if value >= start && value <= end {
			count++
		}
	}

	return count
}

// 复杂度过高。
//把相同喜好的放在一起，然后看在不在区间里面
