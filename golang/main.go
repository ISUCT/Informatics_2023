package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab4/lab5"
)

func main() {
	myslicefortask2 := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	fmt.Println("Кузьмин Владислав Иванович")
	fmt.Println(lab4.Primer1(0.26, 0.66, 0.08))
	fmt.Println(lab4.Primer2(myslicefortask2))

	//lab5

	mouse, err := lab5.NewMouse(5, 5, 5, "Mickey Mouse")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Привет, друг! Вот характеристика мыши:")
	fmt.Printf("Вес: %.1f\n", mouse.GetWeight())
	fmt.Printf("Рост: %.1f\n", mouse.GetHeight())
	fmt.Printf("Возраст: %d\n", mouse.GetAge())
	fmt.Printf("Имя: %s\n", mouse.GetName())

	err = mouse.SetWeight(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = mouse.SetHeight(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = mouse.SetAge(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	mouse.SetName("MiniMouse")
	fmt.Println("Обновленные значения:")
	fmt.Println("Привет, друг! Вот характеристика мыши:")
	fmt.Printf("Вес: %.1f\n", mouse.GetWeight())
	fmt.Printf("Рост: %.1f\n", mouse.GetHeight())
	fmt.Printf("Возраст: %d\n", mouse.GetAge())
	fmt.Printf("Имя: %s\n", mouse.GetName())
}
