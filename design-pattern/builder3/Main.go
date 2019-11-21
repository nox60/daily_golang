package main

import (
	"./builder"
	"fmt"
)

func main() {
	c := builder.Test{Name: "aa"}
	fmt.Println(c)
	fmt.Println("-------")
	vehi := builder.VehicleProduct{}
	builder := builder.ManufacturingDirector{}
	builder.SetBuilder(vehi)

}
