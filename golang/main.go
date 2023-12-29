package main

import (
	"fmt"
	"log"

	"github.com/stretchr/testify/internal/lab4"
	"github.com/stretchr/testify/internal/pistolStruct"
)

func checkErrors(err error) {
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
	fmt.Println(lab4.TaskA(xBeg, xEnd, detX))

	fmt.Println("Task B:")
	var list = []float64{0.2, 0.3, 0.38, 0.43, 0.57}
	fmt.Println(lab4.TaskB(list))

	fmt.Println("Lab5:")
	pistol, err := pistolStruct.NewPistol(9, "Desert Eagler", ".357 Magnum")
	checkErrors(err)

	pistol2, err := pistolStruct.NewPistol(15, "...", "...")
	checkErrors(err)

	pistol2.SetModel("Glock")
	pistol.SetModel("Revolver")

	err = pistol.SetmagazineCapacity(6)
	checkErrors(err)
	fmt.Printf("Pistol's magazine capacity is %d\n", pistol.GetmagazineCapacity())
	fmt.Printf("Its model is %s\n", pistol.GetModel())

	fmt.Printf("Pistol's magazine capacity is %d\n", pistol2.GetmagazineCapacity())
	fmt.Printf("Its model is %s\n", pistol2.GetModel())
}
