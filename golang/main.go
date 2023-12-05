package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/structure"
)

func main() {
	var Toyota structure.Car = structure.Car{Name: "Toyota", Color: "White", Weight: 1500, MaxSpeedKMPH: 250}
	var RiverBoats structure.Boat = structure.Boat{Name: "RiverBoats", LiftingCapacity: 100, MaxSpeedKnots: 10}
	fmt.Printf("Вингорадов Антон Алексеевич 1/278 \n")
	fmt.Println(internal.Cycle(1.25, 3.25, 0.4))
	fmt.Println(internal.Massive([]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
	fmt.Printf("\nThe color of the car was changed from %s \n", Toyota.ChangeColorCar("Black"))
	Toyota.Move()
	Toyota.TableCar()
	vehicels := []structure.VehicleID{Toyota, RiverBoats}
	structure.Drive(vehicels)
}
