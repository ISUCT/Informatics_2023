package country

import (
	"fmt"
)

// Определение структуры страны
type Country struct {
	CountryName string
	Population  int
	Area        float64
	Capital     string
}

// Методы для структуры Country
func (c Country) GetPopulation() int {
	return c.Population
}

func (c Country) GetArea() float64 {
	return c.Area
}
func (c Country) GetCapital() string {
	return c.Capital
}
func (c Country) GetCountryName() string {
	return c.CountryName
}

func main() {
	// Список стран для выбора
	countries := map[string]Country{
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
