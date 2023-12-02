package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/catStruct"
	"isuct.ru/informatics2022/internal/lab4"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printResult(xL []float64, yL []float64) {
	for i := range yL {
		fmt.Printf("x = %.2f, y = %f\n", xL[i], yL[i])
	}
}

func main() {
	fmt.Println("Аксенов Иван Дмитриевич")
	fmt.Println("---")

	var b = 2.5
	var begX = 1.28
	var finX = 3.28
	var detX = 0.4

	var yList, corrXList = lab4.TaskA(begX, finX, detX, b)
	printResult(corrXList, yList)

	var xList = []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	yList = lab4.TaskB(xList, b)
	printResult(xList, yList)
	fmt.Println("---")

	cat, err := catStruct.NewCat(6, "Tiger", "Siamese")
	checkForError(err)
	anotherCat, err := catStruct.NewCat(-1, "No name", "unknown")
	checkForError(err)

	cat.SetAge(3)
	checkForError(err)

	anotherCat.SetName("Marshal")
	cat.SetName("Luna")

	fmt.Printf("Cat's age is %d\n", cat.GetAge())
	fmt.Printf("Its name is %s\n", cat.GetName())
	cat.PetTheCat()

	fmt.Printf("Another cat's age is %d\n", anotherCat.GetAge())
	fmt.Printf("Its breed is %s\n", anotherCat.GetBreed())
	anotherCat.PetTheCat()
}
