package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Орлова Екатерина Евгеньевна, 1/278")
	fmt.Println(internal.Task_A(1.5, 3.5, 0.4))
	fmt.Println(internal.Task_B([]float64{1.9, 2.15, 2.34, 2.74, 3.16}))

	c := internal.City{Country: "Россия", Population: 360687, NameCity: "Иваново"}
	c.SetCountry("Норвегия")
	fmt.Println(c.GetCountry())
	c.Output()
}
