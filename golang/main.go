package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal/struct2"
)

func main() {
	fmt.Println("Полозков Максим Михайлович")
	a := 2.0
	x_n := 1.2
	x_k := 4.2
	x_z := 0.6
	for x := x_n; x <= x_k; x = x + x_z {
		var ya = (math.Pow(math.Log10(a+x), 2) / math.Pow(a+x, 2))
		fmt.Println("Результат задачи A =", ya)
	}
	var argument = [5]float64{1.16, 1.32, 1.47, 1.65, 1.93}
	for z := 0; z < 5; z++ {
		var yb = (math.Pow(math.Log10(a+argument[z]), 2) / math.Pow(a+argument[z], 2))
		fmt.Println("Результат задачи B =", yb)
	}

	m := struct2.Mouse{Brand: "Logitech", Buttons: 3, Wireless: true}

	brand, err := m.GetBrand()
	if err != nil {
		fmt.Println("Ошибка при получении бренда мыши:", err)
	} else {
		fmt.Println("Бренд мыши:", brand)
	}

	if err := m.SetButtons(5); err != nil {
		fmt.Println("Ошибка при установке количества кнопок:", err)
	} else {
		fmt.Println("Количество кнопок после установки:", m.Buttons)
	}
	fmt.Println("Беспроводная мышь?", m.IsWireless())
}
