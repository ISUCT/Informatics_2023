package main

import (
	"fmt"

	"isuct.ru/informatics2022/character"
	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Кувшинов Владислав")
	fmt.Println(internal.CalculateFunctionInRanges(0.26, 0.66, 0.08))
	fmt.Println(internal.CalculateFunctionWithXValues([]float64{0.1, 0.35, 0.4, 0.55, 0.6}))

	player := character.Character{Name: "Гарри Дюбуа", Intellect: 8, Psyche: 7, Physique: 6, Motorics: 5}

	player.PrintDetails()

	// Изменение имени персонажа
	player.SetName("Диджей Миша")
	fmt.Println("Имя персонажа после изменения:", player.Name)

	// Вывод общей силы персонажа
	power := player.GetOverallPoints()
	fmt.Println("Общая сила персонажа:", power)
}
