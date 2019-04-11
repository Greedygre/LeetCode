package 题目

import (
	"fmt"
	"time"
)

// Q: 接口
//
type Queue struct {
	TaskChan chan *Task
}
type Task struct {
	code int
}

type T1 interface {
	RealResponse() interface{}
}

func (q *Queue) RealResponse() {
	fmt.Println(<-q.TaskChan)
}
func (q *Queue) submit(task *Task) {
	select {
	case q.TaskChan <- task:
		fmt.Println("chan task:", task.code)
	case <-time.After(5000 * time.Millisecond):
		fmt.Println("lulala")

	}
}

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

// 定义了有两个方法的接口 I，结构 S 实现了此接口
type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}
func main() {
	//t1 := &Task{
	//	code: 5,
	//}
	//chanT := make(chan *Task, 2)
	//q := &Queue{
	//	TaskChan: chanT,
	//}
	//q.submit(t1)
	//q.submit(&Task{code: 100})
	//q.RealResponse()
	test := &S{
		i: 66,
	}
	f(test)
}

// 结构图体S实现了get和put，定义方法f时候，可以直接写接口，然后调用所有实现了接口的方法。多态/

//
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//type U struct {
//	g int32
//}
//
//// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
//func (t T) M() {
//	fmt.Println(t.S)
//}
//func (u U) M() {
//	fmt.Println(u.g)
//}
//
//
//func main() {
//	var i I = T{"hello"}
//	var ii I =U{888}
//	i.M()
//	ii.M()
//}
