package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	var Toyota internal.Car = internal.Car{Name: "Toyota", Color: "White", Weight: 1500, MaxSpeedKMPH: 250}
	fmt.Printf("Вингорадов Антон Алексеевич 1/278")
	fmt.Println(internal.Cycle(1.25, 3.25, 0.4))
	fmt.Println(internal.Massive([]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
	fmt.Printf("The color of the car was changed from %s \n", Toyota.ChangeColor("Black"))
	Toyota.Move()
	Toyota.Table()
}
