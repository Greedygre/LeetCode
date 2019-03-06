package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	m := 0
	fmt.Scan(&m)
	c := 0
	// 手串上的颜色一共有c种
	fmt.Scan(&c)

	//map[color][]出现的珠子
	color := make(map[int][]int)
	//输入一串字符切割。用于一行行输入的时候
	//part := strings.Split(pree, " ")
	//		for _, pre := range part {
	for i := 0; i < n; i++ {
		pree := 0
		fmt.Scan(&pree)
		if pree != 0 {

			in := 0
			for j := 0; j < pree; j++ {

				fmt.Scan(&in)
				if _, ok := color[in]; ok {
					//存在
					cur := color[in]
					cur = append(cur, i+1)
					color[in] = cur
					//fmt.Println(in, "颜色存在，i：", i+1)
					continue
				}
				//如果不存在
				color[in] = []int{i + 1}
				//fmt.Println("颜色", in, "不存在,i:", i+1)
			}
		}
	}

	fmt.Println(search(n, m, c, color))

}

func search(n, m, c int, color map[int][]int) int {

	count := 0
	for i := 1; i < c+1; i++ {
		colorpart := color[i]
		if n < 2*m {
			if len(colorpart) >= 2 {
				count++
				continue
			}
		}
		//sort.Ints(colorpart)
		if len(colorpart) != 1 {
			//fmt.Println(i, "i颜色不单一的，出现的地方：", colorpart)
			for j, present := range colorpart {
				if j != len(colorpart)-1 && colorpart[j+1]-m < present {
					//fmt.Println(present, "present")
					count++
					break
				}
				//fmt.Println("present", present, "j", j)
				if n-present <= m {
					//fmt.Println("test========")
					if j != len(colorpart)-1 && colorpart[j+1]-m < present {
						count++
						break
					}
					x := m + present - n - 1
					//x应该是9
					//fmt.Println("test")
					for hh := 0; hh < len(colorpart); hh++ {
						//fmt.Println("tesssss", "x:", x, "m:", m, "present:", present)
						if colorpart[hh] >= 1 && colorpart[hh] < x {
							//fmt.Println("te9999999")
							count++
							break
						}
					}

				}
				//fmt.Println("color:", present, "end")
			}
		}

	}

	return count
}

//把相同喜好的放在一起，然后看在不在区间里面
