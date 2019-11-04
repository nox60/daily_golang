package main

import "fmt"

//定义接口
type TypeCalculator interface {
	TypeCal() string
}

type Worker struct {
	Type int
	Name string
}

type Student struct {
	Name string
}

type EmptyStr struct {
}

//对接口的实现1
func (w Worker) TypeCal() string {
	if w.Type == 0 {
		return w.Name + "是蓝翔毕业的员工"
	} else {
		return w.Name + "不是蓝翔毕业的员工"
	}
}

//对接口的实现2
func (s Student) TypeCal() string {
	return s.Name + "还在蓝翔学挖掘机炒菜"
}

//对接口的实现3
func (*EmptyStr) TypeCal() string {
	return "还没有考上蓝翔"
}

func main() {
	worker := Worker{Type: 0, Name: "小华"}
	student := Student{Name: "小明"}
	empty := EmptyStr{}
	workers := []TypeCalculator{worker, student, &empty}
	for _, v := range workers {
		fmt.Println(v.TypeCal())
	}

	empty.TypeCal()

	s := SimpleImpl{}

	s.SimpleMethod()
}

type NoName interface {
	SimpleMethod() string
}

type SimpleImpl struct {
}

func (s *SimpleImpl) SimpleMethod() string {
	fmt.Println("--->")
	return "done"
}
