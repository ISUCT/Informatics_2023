package main

import (
	"fmt"

	lab4 "command-line-argumentsC:\\Users\\dmitr\\Desktop\\репризиторий\\Informatics_2023\\golang\\internal\\lab 4\\calculate.go"
	lab4 "isuct.ru/informatics2022/internal/lab 4"
)

// вариант 12
func main() {
	var a float64 = 1.6
	fmt.Println("Задача A")
	x, y := lab4.TuskA(a, 1.2, 3.7, 0.5)
	for i, _ := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("Задача B")
	var listik [5]float64 = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	x, y = lab4.TuskB(a, listik)
	for i, _ := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
}
