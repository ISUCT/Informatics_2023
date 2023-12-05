package structure

import "fmt"

type Boat struct {
	Name            string
	LiftingCapacity int
	MaxSpeedKnots   float64
}

func (boat *Boat) Flooding(cargoWeight int) {
	if cargoWeight > boat.LiftingCapacity {
		fmt.Printf("The %s sank due to too heavy cargo", boat.Name)
	}
}
func (boat Boat) Move() {
	fmt.Printf("The %s is moving at a speed of %0.f Knots per hour \n", boat.Name, boat.MaxSpeedKnots)
}
