package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := &a
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
