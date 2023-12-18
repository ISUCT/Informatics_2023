package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	result1 := lab4.Assignment_1(0.8, 0.4, 1.23, 7.23, 1.2)
	result2 := lab4.Assignment_2(0.8, 0.4, []float64{1.88, 2.26, 3.84, 4.55, -6.21})
	fmt.Println(result1)
	fmt.Println(result2)
}
