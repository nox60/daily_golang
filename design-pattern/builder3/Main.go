package main

import (
	"daily_golang/src"
	"fmt"
)

func main() {
	c := src.Test{Name: "aa"}
	fmt.Println(c)
	fmt.Println("-------")
	vehi := src.VehicleProduct{}
	//builder := builder.ManufacturingDirector{}
	carbuilder := src.CarBuilder{}
	carbuilder.VehicleProduct = vehi
	car := carbuilder.SetSeats(4).SetStructure("ss").SetWheels(4).GetVehicle()
	fmt.Println(car)
}
