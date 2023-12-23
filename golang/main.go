package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab3"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	a1 := lab3.Task1(1.2, 0.48, 0.7, 2.2, 0.3)
	a2 := lab3.Task2(1.2, 0.48, []float64{0.25, 0.36, 0.56, 0.94, 1.28})
	fmt.Println(a1)
	fmt.Println(a2)
	
	bunny, err := lab5.NewBunny(10, 4000, "Bunny")
	
	if (err == nil) {
		fmt.Printf("Bunny's age is %d\n", bunny.GetAge())
		fmt.Printf("Bunny's weight is %d\n", bunny.GetWeight())
		fmt.Printf("Bunny's name is %s\n", bunny.GetName())
	} else {
		fmt.Println(err)
	}
}