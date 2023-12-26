package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/cat"
)

func main() {
	fmt.Println("Алексеев Максим Алексеевич | Вариант 1")
	fmt.Println(internal.Zadacha_A(1.28, 3.28, 0.4))
	fmt.Println(internal.Zadacha_B([]float64{1.1, 2.4, 3.6, 1.7, 3.9}))

	cat1, err := cat.NewCat("Барсик", 10, "black")
	if err != nil {
		fmt.Printf("%v", err)
	}
	cat1.Meow()
	fmt.Printf("\n%v\n", cat1.GetColor())
	cat1.SetColor("white")
	fmt.Printf("%v\n", cat1.GetColor())
	fmt.Printf("%v\n", cat1.GetAge())
	cat1.SetAge(30)
	fmt.Printf("%v\n", cat1.GetAge())
	cat2, err := cat.NewCat("Кузя", 77, "orange")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(cat2.GetName())
	fmt.Println(cat2.GetAge())
}
