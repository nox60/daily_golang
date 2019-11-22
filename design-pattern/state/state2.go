package main

import "fmt"

type StateContext struct {
}

type State interface {
	WriteName(name string)
}

func main() {
	fmt.Println("")
}
