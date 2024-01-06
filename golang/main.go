package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	fmt.Println("Сысуйкин Денис Дмитриевич")
	SliceTaskA := lab4.TaskA(2.25, 1.2, 2.7, 0.3)
	fmt.Println("Результаты TaskA:", SliceTaskA)
	SliceForTaskB := []float64{2.25, 1.31, 1.39, 1.44, 1.56, 1.92}
	resultS := lab4.TaskB(SliceForTaskB, 2.25)
	fmt.Println("Результаты TaskB:", resultS)
	// LAB 5
	fmt.Printf("\nLABORATORY: 5\n")
	cat, err := lab5.NewCat(5, 3.5, 25.6, "Stesha")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Возраст: %d\n", cat.GetAge())
	fmt.Printf("Вес: %.2f\n", cat.GetWeight())
	fmt.Printf("Рост: %.3f\n", cat.GetHeight())
	fmt.Println("Имя:", cat.GetName())
}
