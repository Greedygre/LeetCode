package main

import "fmt"

func test3()  {
	n := 0
	fmt.Scan(&n)
	wulizhi := []int{}
	jinbi := []int{}
	for i:=0;i<n;i++{
		pre := 0
		fmt.Scan(&pre)
		wulizhi = append(wulizhi,pre)
	}
	for i:=0;i<n;i++{
		pre := 0
		fmt.Scan(&pre)
		jinbi = append(jinbi,pre)
	}
	//fmt.Println(jinbi,wulizhi)
	cost :=0
	max := 0
	for i:=0;i<n;i++{
		if i==0{
			cost += jinbi[i]
			max = wulizhi[i]
			continue
		}
		if max < wulizhi[i]{
			cost+=jinbi[i]
			max += wulizhi[i]
			continue
		}
		for j:=i;j<n;j++{
			if wulizhi[j]>max{
				curcost,jiawulizhi := finddashou(wulizhi[i:j],jinbi[i:j],wulizhi[j],max)
				max += jiawulizhi
				cost+=curcost
				break
			}
		}

	}
	fmt.Println(cost)
		}


func finddashou(wulizhi,jinbi []int,tobe int,max int) (int,int) {

cur := 0
for i,value := range  wulizhi{
	if max + value >= tobe {
		cur = jinbi[i]
		max = value
		break
	}else {
		cur += jinbi[i]
		max += value
	}
	}
return cur,max
}
func main()  {
	test3()
}