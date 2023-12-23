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
	lisitca1, _ := lab5.NewAnimalLisitca(
		"Alisa",
		"orange",
		4,
	)
	fmt.Println(lisitca1)
	fmt.Println(lisitca1.Voice("шурхшурхшурх"))
	fmt.Println(lisitca1.Calling())
}
