package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Долев Максим Алексеевич")
	fmt.Println("Lab 4")

	a := 2.0
	b := 0.95
	x_beg := 1.25
	x_end := 2.75
	x_stp := 0.3

	xList := []float64{2.2, 3.78, 4.51, 6.58, 1.2}

	fmt.Println(lab4.Task_A(x_beg, x_end, x_stp, a, b))
	fmt.Println(lab4.Task_B(xList, a, b))

	fmt.Println("Lab 5")

	dog, err := lab5.TheDog(8, "Zhuchka", "female")
	checkForError(err)

	fmt.Printf("The dog is %d years old\n", dog.GetAge())
	fmt.Printf("The dog's name is %s\n", dog.GetName())
	fmt.Printf("The dog is a %s, it's easy to check", dog.GetGender())
}
