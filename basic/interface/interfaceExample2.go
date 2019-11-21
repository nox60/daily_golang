package main

import "fmt"

//定义接口
type Interface1 interface {
	FirstMethod() string
	SecondMethod() bool
}

type AnotherInterface interface {
	AnotherMethod() bool
}

type ThirdInterface interface {
	FirstMethod() string
}

type Implement1 struct {
	Type int
	Name string
}

type Implement2 struct {
	Name string
}

type Implement3 struct {
}

//对接口的实现1
func (w Implement1) FirstMethod() string {
	return ""
}

func (w Implement1) AnotherMethod() bool {
	return true
}

//对接口的实现2
func (s Implement2) FirstMethod() string {
	return s.Name + "还在蓝翔学挖掘机炒菜"
}

func (w Implement2) AnotherMethod() bool {
	return true
}

//对接口的实现3
func (*Implement3) FirstMethod() string {
	return "还没有考上蓝翔"
}

func (*Implement3) SecondMethod() bool {
	return false
}

func (w Implement3) AnotherMethod() bool {
	return true
}

func main() {
	implement1 := Implement1{Type: 0, Name: "小华"}
	implement2 := Implement2{Name: "小明"}
	implement3 := Implement3{}
	//workers := []Interface1{implement1, implement2, &implement3}
	//上面这行代码要报错，因为，Implement1和Implement2并没有实现 SecondMethod, 只有把一个接口里定义完的所有方法全部实现了，才能算是对该接口的完整实现。

	workers := []AnotherInterface{implement1, implement2, &implement3}

	thirdWorkers := []ThirdInterface{implement1, implement2, &implement3}

	for _, v := range workers {
		fmt.Println(v.AnotherMethod())
	}

	for _, v := range thirdWorkers {
		fmt.Println(v.FirstMethod())
	}

}
