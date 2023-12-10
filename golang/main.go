package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Вераксо Марк Владимирович 1/280")
	fmt.Println("1 Вариант")
	fmt.Println(internal.Cycle(1.28, 3.28, 0.4))
	fmt.Println(internal.Srez([]float64{1.1, 2.4, 3.6, 1.7, 3.9}))
	
}
