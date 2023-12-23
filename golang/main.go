package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/structure"
)

func main() {
	Toyota := structure.CreateCar("Toyota", "White", 1500, 250)
	RiverBoats := structure.CreateBoat("River Boats", 100, 100)
	fmt.Printf("Вингорадов Антон Алексеевич 1/278 \n")
	fmt.Println(internal.Cycle(1.25, 3.25, 0.4))
	fmt.Println(internal.Massive([]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
	fmt.Printf("\nThe color of the car was changed from %s \n", Toyota.ChangeColorCar("Black"))
	Toyota.Move()
	Toyota.TableCar()
	vehicels := []structure.VehicleID{Toyota, RiverBoats}
	for i := 0; i != len(vehicels); i++ {
		vehicels[i].Move()
	}
	structure.Drive(vehicels)
}
