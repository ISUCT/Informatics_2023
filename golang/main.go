package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/plane"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Иванов Алексей Алексеевич 1/278")
	fmt.Println(internal.SolveA(1.1, 0.09, 1.2, 2.2, 0.2))
	fmt.Println(internal.SolveB([]float64{4.48, 3.56, 2.78, 5.28, 3.21}))

	fmt.Println("\n First Plane: \n")
	plane1, err := plane.NewPlane("Космолёт", "Фиолетовый", 3124)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(plane1.GetModel())
	fmt.Println(plane1.GetColor())
	plane1.SetColor("Белый")
	fmt.Println(plane1.GetColor())
	fmt.Println(plane1.GetSpeed())
	plane1.SetSpeed(3400)
	fmt.Println(plane1.GetSpeed())

	fmt.Println("\n Second Plane: \n")
	plane2, err := plane.NewPlane("Boeing 737", "Серый", 400)
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Println(plane2.GetModel())
	fmt.Println(plane2.GetColor())
	plane2.SetColor("Красный")
	fmt.Println(plane2.GetColor())
}
