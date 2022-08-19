package main

import "fmt"

type Person interface {
	setName(string)
	getName() string
}
type Student struct {
	name string
	age  int
}

func (stu *Student) setName(name string) {
	println(name)
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student{
		name: "charles",
		age:  18,
	}
	p.setName("xiaowei")
	fmt.Println(p.getName())

}
