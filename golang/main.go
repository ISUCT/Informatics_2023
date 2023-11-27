package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/internal/LAB4"
)

func main() {
	fmt.Println("Гвоздарёв Ярослав Владимирович")
	fmt.Println("задача A")
	xNach := 3.5
	xKon := 6.5
	shag := 0.6
	x, y := lab4.TaskA(xNach, xKon, shag)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("задача B")
	xArr := []float64{2.89, 3.54, 5.21, 6.28, 3.48}
	x, y = lab4.TaskB(xArr)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}

}
