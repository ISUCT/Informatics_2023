package main

import (
	"fmt"
	"log"

	"github.com/stretchr/testify/internal/lab4"
	"github.com/stretchr/testify/internal/pistolStruct"
)

func result(x []float64, y []float64) {
	for i := range y {
		fmt.Printf("x = %.2f, y = %f\n", x[i], y[i])
	}
}

func inputErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Хосровян Александр Арменович")

	xBeg := 0.11
	xEnd := 0.36
	detX := 0.05

	fmt.Println("Task A:")
	var yList, xlist2 = lab4.TaskA(xBeg, xEnd, detX)
	result(xlist2, yList)

	fmt.Println("Task B:")
	var xList = []float64{0.2, 0.3, 0.38, 0.43, 0.57}

	yList = lab4.TaskB(xList)
	result(xList, yList)

	fmt.Println("Lab5:")
	var pistol = pistolStruct.NewPistol(9, "Desert Eagler", ".357 Magnum")

	var err error = pistol.SetmagazineCapacity(20)
	inputErrors(err)

	pistol.SetModel("Glock 18")
	pistol.SetCaliber("9х19 Luger")

	fmt.Printf("Pistol's magazineCapacity is %d\n", pistol.GetmagazineCapacity())
	fmt.Printf("Its model is %s\n", pistol.GetModel())
	fmt.Printf("Caliber of this pistol is %s\n", pistol.GetCaliber())
	pistol.PulltheTriger()
}
