package structure

import "fmt"

type Boat struct {
	name            string
	liftingCapacity int
	maxSpeedKnots   float64
}

func (boat *Boat) Flooding(cargoWeight int) {
	if cargoWeight > boat.liftingCapacity {
		fmt.Printf("The %s sank due to too heavy cargo", boat.name)
	}
}
func (boat Boat) Move() {
	fmt.Printf("The %s is moving at a speed of %0.f Knots per hour \n", boat.name, boat.maxSpeedKnots)
}
func CreateBoat(name string, lift int, speed float64) *Boat {
	boat := new(Boat)
	boat.name = name
	boat.liftingCapacity = lift
	boat.maxSpeedKnots = speed
	return boat
}
