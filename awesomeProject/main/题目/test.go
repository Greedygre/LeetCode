package main

import "fmt"

func main() {
	ma := make(map[string]bool, 4)
	ma["小明"] = true
	ma["xiaohong"] = true
	for m := range ma {
		fmt.Println(ma[m])
	}
}
