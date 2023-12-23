package structure

import "fmt"

type Car struct {
	name         string
	color        string
	weight       float64
	maxSpeedKMPH float64
}

func (car *Car) ChangeColorCar(RequiredColor string) string {
	OldColor := car.color
	car.color = RequiredColor
	changes := OldColor + " to " + car.color
	return changes
}
func (car Car) Move() {
	fmt.Printf("The %s is moving at a speed of %0.f Kilometers per hour \n", car.name, car.maxSpeedKMPH)
}

func (car *Car) TableCar() {
	if car.weight > 100 {
		fmt.Printf("Your table is broken because the %0.f kilogram machine is too heavy \n \n", car.weight)
	}
}
func CreateCar(name string, color string, wght float64, speed float64) *Car {
	car := new(Car)
	car.name = name
	car.color = color
	car.weight = wght
	car.maxSpeedKMPH = speed
	return car
}
