package 题目

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// 有一个首尾相接的字符序列，全部元素已大写字符表示，要求截取一段包含ABCDE的连续子串，返回这一子串的长度，例如ABCYDYE，返回6，ATTMBQECPD返回7
// 求包含abcd的最小长度
func find(str string) int {
	// 一前一后两个指针
	return 0
}

func main() {
	const jsonStream = `
	{"Name": "Ed", 
"Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var m Message
	if err := dec.Decode(&m); err == io.EOF {

	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s\n", m.Name, m.Text)

}
