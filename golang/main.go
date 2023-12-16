package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	var venti internal.Character = internal.Character{Name: "Venti", Class: "Bard", Level: 1, HP: 10}
	fmt.Println("Красильникова Анастасия 1/278")
	fmt.Println(internal.Cycle(1.81, 5.31, 0.7))
	fmt.Println(internal.Array([]float64{2.4, 2.8, 3.9, 4.7, 3.16}))
	fmt.Println("Level up", venti.LevelUp(venti.Level))
	fmt.Printf("HP up. Current value: %d \n", venti.HpChange(venti.Level))
	fmt.Printf("Chatacter class was changed to %s", venti.ClassChange("Wizard"))
}
