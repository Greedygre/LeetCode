package main

import (
	"fmt"
	"time"
)

type Duck interface {
	GaGaSpeaking()
	OfficialWalking()
}
type Actor interface {
	MakeFun()
}
type DonaldDuck struct {
	height uint
	name   string
}

func (dd *DonaldDuck) GaGaSpeaking()    { fmt.Println("DonaldDuck gaga") }
func (dd *DonaldDuck) OfficialWalking() { fmt.Println("DonaldDuck walk") }
func (dd *DonaldDuck) MakeFun()         { fmt.Println("DonaldDuck make fun") }
func main2() {
	dd := &DonaldDuck{10, "tang lao ya"}
	var duck Duck = dd
	var actor Actor = dd
	duck.GaGaSpeaking()
	actor.MakeFun()
	dd.OfficialWalking()
}
func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(5 * time.Second)
	//ch := make(chan string, 2)
	i := 0
	for {
		i++
		select {
		case <-tick:
			fmt.Println(tick, "lullal")
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}

	}

	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x) // "0" "2" "4" "6" "8"
	//	case ch <- i:
	//	}
	//}
}
