package 题目

import "fmt"

// 思路：统计出现次数；维护一个二维数组，每一行
//a字符串中出现的次数
//b
func commonChars(A []string) []string {

	m := make(map[int32]int)
	for _, str := range A {
		cur := make(map[int32]int)
		for _, cor := range str {
			if _, ok := cur[cor]; ok {
				// 记录前面出现的最小的
				cur[cor]++
			} else {
				cur[cor] = 1
			}
		}

	}
}
func main() {
	fmt.Println('5')
}
