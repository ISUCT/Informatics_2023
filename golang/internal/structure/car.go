package structure

import "fmt"

type Car struct {
	Name         string
	Color        string
	Weight       float64
	MaxSpeedKMPH float64
}

func (car *Car) ChangeColorCar(RequiredColor string) string {
	OldColor := car.Color
	car.Color = RequiredColor
	changes := OldColor + " to " + car.Color
	return changes
}
func (car Car) Move() {
	fmt.Printf("The %s is moving at a speed of %0.f Kilometers per hour \n", car.Name, car.MaxSpeedKMPH)
}

func (car *Car) TableCar() {
	if car.Weight > 100 {
		fmt.Printf("Your table is broken because the %0.f kilogram machine is too heavy \n \n", car.Weight)
	}
}
