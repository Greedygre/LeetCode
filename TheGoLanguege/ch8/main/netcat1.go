// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)
var wg sync.WaitGroup
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// os.Stderr是红色的标准输出字体
	wg.Add(2)
	go mustCopy(os.Stderr, conn)
	// os.Stdout是白色的
	 mustCopy(os.Stdout,conn)
	wg.Wait()
}

func mustCopy(dst io.Writer, src io.Reader) {
	defer wg.Done()
	if written, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}else {
		fmt.Println(written)
	}
fmt.Println("lulalal")
}