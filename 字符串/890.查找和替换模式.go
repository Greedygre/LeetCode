package main

import (
	"fmt"
)


func findAndReplacePattern(words []string, pattern string) []string {
result := []string{}
for i:=0;i<len(words);i++{
	count1 := make(map[uint8]uint8)
	if len(words[i])!=len(pattern){
		continue
	}
	count:=make(map[uint8]uint8)
	for i:=0;i<len(pattern);i++{
		if _,ok:=count[pattern[i]];ok{
			continue
		}else {
			count[pattern[i]]=1
		}
	}
	lenp:=len(count)
	for j:=0;j<len(words[i]);j++{
		//如果已经映射过了
		if value,ok:=count1[words[i][j]];ok{
			//fmt.Println("===已经匹配过")
			if value!=pattern[j]{
				//fmt.Println("===已经匹配过，不匹配")
				break
			}
		}else {
			count1[words[i][j]] = pattern[j]
			fmt.Println(pattern[j])
		}
		if j == len(words[i])-1&&len(count1)==lenp{
			//fmt.Println("============w",words[i])
			result = append(result, words[i])
		}
	}

}
return result
}


func main()  {
	fmt.Println('a','z','A','Z')//97 65
	words := []string{"abc","deq","mee","aqq","dkd","ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words,pattern))
}