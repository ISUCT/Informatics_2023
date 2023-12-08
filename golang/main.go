package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

const a float64 = 2.0
const b float64 = 0.95
const xStart float64 = 1.25
const xEnd float64 = 2.75
const xStep float64 = 0.3

func main() {
	fmt.Println("Быстров Владислав Евегньевич")

	fmt.Println("Задача А:")
	stepsResults := internal.CalculateFuncForSteps(a, b, xStart, xEnd, xStep)
	fmt.Println(stepsResults)

	fmt.Println("Задача Б:")
	xesResults := internal.CalculateFuncForXes(a, b, []float64{2.2, 3.78, 4.51, 6.58, 1.2})
	fmt.Println(xesResults)
}
