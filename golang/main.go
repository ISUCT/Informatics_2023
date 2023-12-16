package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	fmt.Printf("Jakovleva Darja Dmitrievna 1/278\n\n")
	fmt.Println("24 вариант")
	fmt.Println("Task_A:", lab4.Task_A(7.2, 1.3, 1.56, 4.71, 0.63))
	fmt.Println("Task_B:", lab4.Task_B(7.2, 1.3, []float64{2.4, 2.8, 3.9, 4.7, 3.16}))

	fmt.Println("2 вариант: лисица")
	var lisa1 lab5.Lisica = lab5.Lisica{
		Name:  "Alisa",
		Color: "orange",
		Age:   4,
	}
	fmt.Println(lisa1.GetAge())
	fmt.Printf("The color of the %s is %s \n", lisa1.Name, lisa1.GetColor())
}
