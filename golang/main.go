package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Вингорадов Антон Алексеевич 1/278")
	fmt.Println(internal.Cycle(1.25, 3.25, 0.4))
	fmt.Println(internal.Massive([5]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
}
