package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/internal"
)

func main() {

	lab4.Zad1()

	values := []float64{0.1, 0.3, 0.4, 0.45, 0.65}

	result := lab4.Zad2(values)

	for i, x := range values {
		formattedResult := fmt.Sprintf("%.3f", result[i])
		fmt.Println("При x=", x, " y=", formattedResult)
	}
	fmt.Println("Мазуров Даниил Алексеевич")

}
