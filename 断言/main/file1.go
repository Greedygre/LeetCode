package aa

import (
	"fmt"
	"断言/aa"
	//"reflect"
)

type c struct {
	name string
}

func Test(b tester) {
	if a, ok := b.(c); ok {
		a.d()
	}
}

type tester interface {
	d()
}

func (c c) d() {
	fmt.Println(c.name)
}
func main2() {
	var t tester
	t = c{"唐老鸭"}
	Test(t)
	aa.Test2()
}

type Tester interface {
	getName() string
}
type Tester2 interface {
	printName()
}

//===Person类型====
type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}
func (p Person) printName() {
	fmt.Println(p.name)
}

//============

func check(t Tester) {
	//第一种情况
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("%T\n%s\n", f, f.getName())
	}
	//第二种情况
	if t, ok2 := t.(Tester2); ok2 { //重用变量名t（无需重新声明）
		check2(t) //若类型断言为true，则新的t被转型为Tester2接口类型，但其动态类型和动态值不变
	}
}
func check2(t Tester2) {
	t.printName()
}
