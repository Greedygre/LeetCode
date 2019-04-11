package main

import "sync"
import "log"

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("i:%d", i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("exit")
}
