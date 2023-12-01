package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/laba4"
)

func main() {
	fmt.Println("Бусыгин Андрей Михайлович")
	a1 := laba4.TaskA(1.25, 3.25, 0.4)
	a2 := laba4.TaskB([]float64{1.84, 2.71, 3.81, 4.56, 5.62})
	fmt.Println(a1)
	fmt.Println(a2)
}
