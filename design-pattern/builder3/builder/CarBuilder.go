package builder

import "daily_golang/src"

type Test struct {
	Name string
}

type CarBuilder struct {
	VehicleProduct src.VehicleProduct
}

func (c *CarBuilder) SetWheels(wheels int) src.BuildProcess {
	c.VehicleProduct.Wheels = wheels
	return c
}

func (c *CarBuilder) SetSeats(seats int) src.BuildProcess {
	c.VehicleProduct.Seats = seats
	return c
}

func (c *CarBuilder) SetStructure(structure string) src.BuildProcess {
	c.VehicleProduct.Structure = structure
	return c
}

func (c *CarBuilder) GetVehicle() src.VehicleProduct {
	return c.VehicleProduct
}
