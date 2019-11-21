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
	//builder := builder.ManufacturingDirector{}
	carbuilder := builder.CarBuilder{}
	carbuilder.VehicleProduct = vehi
	car := carbuilder.SetSeats(4).SetStructure("ss").SetWheels(4).GetVehicle()

	fmt.Println(car)

}
