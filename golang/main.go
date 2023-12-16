package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	country "isuct.ru/informatics2022/internal/Country"
)

func main() {
	var a = 0.4
	var b = 0.8
	fmt.Println(internal.TaskA(a, b, 3.2, 6.2, 0.6))
	var xList = []float64{4.48, 3.56, 2.78, 5.28, 3.21}
	fmt.Println(internal.TaskB(xList, a, b), xList)
	//lab5
	// Список стран для выбора
	countries := map[string]country.Country{
		"USA": {
			CountryName: "USA",
			Population:  333287557,
			Area:        9629091.0,
			Capital:     "Washington",
		},
		"Russia": {
			CountryName: "Russia",
			Population:  145934462,
			Area:        17098242.0,
			Capital:     "Moscow",
		},
		"China": {
			CountryName: "China",
			Population:  1439323776,
			Area:        9640011.0,
			Capital:     "Beijing",
		},
	}
	// Запрос выбора страны у пользователя
	var selectedContry string
	fmt.Print("Выберете страну из списка:")
	fmt.Scanln(&selectedContry)
	country, ok := countries[selectedContry]
	if !ok {
		fmt.Println("Страна не найдена")
		return
	}

	//Вывод данных о стране
	fmt.Println("Страна:", country.GetCountryName())
	fmt.Println("Популяция:", country.GetPopulation())
	fmt.Println("Площадь:", country.GetArea())
	fmt.Println("Столица:", country.GetCapital())
}
