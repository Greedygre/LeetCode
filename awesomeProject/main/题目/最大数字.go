package 题目

import (
	"fmt"
	"sort"
	"strconv"
)

//思路，重写sort
type element struct {
	code string
}
type elements []element

func smallestNumber(nums []int) string {
	res := elements{}
	//gg := []element{}
	for _, cur := range nums {
		pre := strconv.Itoa(cur)
		res = append(res, element{code: pre})
	}
	sort.Sort(res)
	result := ""
	for _, i := range res {
		if i.code != "0" {
			result = fmt.Sprintf("%s%s", result, i.code)
		}
	}
	if len(result) == 0 {
		return "0"
	}
	return result
}

func (e elements) Len() int {
	return len(e)
}
func (e elements) Less(i, j int) bool {
	str_i := fmt.Sprintf("%s%s", e[i].code, e[j].code)
	str_j := fmt.Sprintf("%s%s", e[j].code, e[i].code)
	//if str_i == str_j {
	//	return true
	//}
	//
	//if str_i[0] > str_j[0] {
	//	return false
	//} else if str_i[0] == str_j[0] {
	//
	//	for ii := 1; ii < len(str_i); ii++ {
	//		if str_i[ii] > str_j[ii] {
	//			return false
	//		} else if str_i[ii] == str_j[ii] {
	//			continue
	//		} else {
	//			return true
	//		}
	//	}
	//} else {
	//	return true
	//}
	return str_i < str_j
}
func (e elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func main() {
	fmt.Println(smallestNumber([]int{12, 12, 121, 0}))

	fmt.Println(smallestNumber([]int{12, 12121, 1211, 12}))
	fmt.Println(smallestNumber([]int{0, 0, 16, 4, 36, 297}))
	fmt.Println(smallestNumber([]int{0, 0, 0, 0, 0}))
}
