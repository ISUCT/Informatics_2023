package structure

type VehicleID interface {
	Move()
}

func Drive(vehicle []VehicleID) {
	for i := 0; i != len(vehicle); i++ {
		vehicleCurrent := vehicle[i]
		vehicleCurrent.Move()
	}
}
