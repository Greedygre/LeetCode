package d

import "fmt"

//
//3. 目录a里面有b.go、c.go 两个文件，go文件第一行 package d ，b.go里面有func E()
// - 现在想用这个E()，怎么用？想用 f.E() 这个句子来用，该怎么用
func E() {
	fmt.Println("lulal")
}
