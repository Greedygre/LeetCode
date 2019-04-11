package main

import (
	"fmt"
	"reflect"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func main() {
	const x = 10
	fmt.Println(*f(x))
}
func print(a ...interface{}) {
	fmt.Println(a...)
	b := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

}
func f(x int) *int {
	return &x
}
func g() int {
	x = new(int)
	return *x
}
