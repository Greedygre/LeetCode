package 题目

import (
	"fmt"
	"strconv"
)

// 9位数字，
// 不能包含5位以上重复，
// 首位不为0，
// 不能包含AAABBBCCC,
// 不能包含五位以上的递增和递减
func count() {
	co := 0
	for i := 100000000; i < 999999999; i++ {
		if no5same(i) && no5zengOrJian(i) {
			co++
		}
	}
	fmt.Println(co)
}
func no5same(num int) bool {
	str := strconv.Itoa(num)
	byt := []byte(str)
	BubbleAsort(byt)
	c := 0
	v1 := byt[0]
	for _, v := range byt {
		if v1 == v {
			c++
			if c >= 5 {
				return false
			}
			continue
		}
		v1 = v
		c = 1
	}
	return true
}
func BubbleAsort(values []byte) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	//fmt.Println(values)
}

//func BubbleZsort(values []int) {
//	for i := 0; i < len(values)-1; i++ {
//		for j := i + 1; j < len(values); j++ {
//			if values[i] < values[j] {
//				values[i], values[j] = values[j], values[i]
//			}
//		}
//	}
//	fmt.Println(values)
//}
func BubbleZsort(values []byte) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	//fmt.Println(values)
}
func no5zengOrJian(num int) bool {
	str := strconv.Itoa(num)
	byt := []byte(str)
	BubbleAsort(byt)
	str1 := string(byt)
	byt = []byte(str)
	BubbleZsort(byt)
	str2 := string(byt)
	for i := 0; i < 5; i++ {
		fmt.Println(str1[str[i]-49:str[i]+5-49], str[i], '0')
		if str[i] < 48 {
			// 57是9
			if str[i:i+5] == str1[str[i]-48:str[i]+5-48] {
				return false
			}

			if str[i:i+5] == str2[str[i]-48:str[i]+5-48] {
				return false
			}
		}

	}
	return true
}
func main() {

	count()
	//fmt.Println(str2[3:8])
	//fmt.Println(str1[:5] == str2[4:9])
}
