package main

import (
	"fmt"
	"sync"
	"time"
)

var result int
var a chan int
var b chan int
var wg sync.WaitGroup

func main() {
	Test2()
}

// 2. 两个管子a,b  两个线程，线程1每0.02秒往a放一个整数，线程2每0.03秒往b放一个整数，
// 线程c的逻辑，延迟2秒启动：var result，如果管子a有数字，result=result+这个数字，
// 如果管子a没有数字的时候，从b取1个数字，result=result-这个数字，两个管子都没有数字的时候，
// 让线程1、2停止生产，停止生产2s后退出。 写出完整可运行的程序
func Test2() {
	// 无缓存
	a = make(chan int, 5)
	b = make(chan int, 5)
	// 线程1
	i := 1
	go func() {
		for {
			i := i * 2
			time.Sleep(20 * time.Millisecond)
			a <- i
		}
	}()
	// 线程2
	go func() {
		for {
			i := i * 3
			time.Sleep(30 * time.Millisecond)
			b <- i
		}
	}()
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("已退出,result:", result)

}
func worker() {
	defer wg.Done()
loop:
	for {
		select {
		case cur := <-a:
			result += cur
			fmt.Println(cur, result, "===")
		default:

			select {
			case cur := <-b:
				result -= cur
				fmt.Println(cur, result)
			default:

				break loop
			}
		}
	}
}
