package main

import (
	"fmt"

	Lab4 "isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Рыбников Юрий Анатольевич")
	fmt.Println("20 лет")
	//Work1
	fmt.Println("AnswerOnWork1:")
	resultWork1 := Lab4.Work1(2.0, 0, 1.2, 4.2, 0.6)
	fmt.Println(resultWork1)
	//Work2
	SliceWork2 := []float64{1.16, 1.32, 1.47, 1.65, 1.93}
	fmt.Println("AnswerOnWork2:")
	resultSlice := Lab4.Work2(SliceWork2, 2.0, 0)
	fmt.Println(resultSlice)
}
