package builder

import "daily_golang/src"

type ManufacturingDirector struct {
	builder src.BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats(4).SetWheels(4).SetStructure("small")
}

func (f *ManufacturingDirector) SetBuilder(b src.BuildProcess) {
	f.builder = b
}
