package builder

type Test struct {
	Name string
}

type CarBuilder struct {
	vehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels(wheels int) BuildProcess {
	c.vehicleProduct.Wheels = wheels
	return c
}

func (c *CarBuilder) SetSeats(seats int) BuildProcess {
	c.vehicleProduct.Seats = seats
	return c
}

func (c *CarBuilder) SetStructure(structure string) BuildProcess {
	c.vehicleProduct.Structure = structure
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicleProduct
}
