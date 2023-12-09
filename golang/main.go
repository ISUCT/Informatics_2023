package main

import (
	"fmt"

	lab4test "isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Жуков Михаил Андреевич")
	fmt.Println("17 лет")
	//task1
	fmt.Println("AnswerOnTask1:")
	resultSliceTaskA := lab4test.Task1Slice(1.1, 0.09, 1.2, 2.2, 0.2)
	fmt.Println(resultSliceTaskA)
	//task2
	mySliceForTaskB := []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	fmt.Println("AnswerOnTask2:")
	resultSlice := lab4test.TaskB(mySliceForTaskB, 1.1, 0.09)
	fmt.Println(resultSlice)
}
