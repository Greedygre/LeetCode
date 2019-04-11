package main

import "fmt"

// 第一行为字符串的长度，第二行为所给字符串
func test2()  {
	n := 0
	fmt.Scan(&n)
	str := ""
	fmt.Scan(&str)
	st := []byte(str)
	str2 := []byte{}
	for i,value := range st {
		//fmt.Println(i,value)
		// copy(str2,st[:i])

			//fmt.Println(str2,"str2")
		if i==0{
			str2 = append(str2,value)
			continue
		}

			if len(str2)!=0 && ((str2[len(str2)-1]==49 &&value==48)||(str2[len(str2)-1]==48 &&value==49) ){
				str2 = str2[:len(str2)-1]
				//fmt.Println(str2)
			}else {
				str2 = append(str2,value)
			}

	}
	fmt.Println(len(str2))
}
func main()  {
test2()
}