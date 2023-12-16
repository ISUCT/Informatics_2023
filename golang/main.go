package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
	answer1 := internal.Cycle(1.28, 3.28, 0.4)
	answer2 := internal.Srez(2.5, []float64{1.1, 2.4, 3.6, 1.7, 3.9})
	fmt.Println("Вераксо Марк Владимирович 1/280")
	fmt.Println("1 Вариант")
	fmt.Println(answer1)
	fmt.Println(answer2)
	
}
