package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	fmt.Println("Lab 5")
	fmt.Println("Телевизор 1")
	tv1, err := lab5.NewTv("Samsung", "Black", 63)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(tv1.GetSound())
	tv1.SetSound(25)
	fmt.Println(tv1.GetSound())
	tv1.TurnOn()
	fmt.Println("\nТелевизор 2")
	tv2, err := lab5.NewTv("Xiaomi", "White", 120)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(tv2.GetModel())
}
