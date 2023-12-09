package main

import (
	"fmt"
	"informatics2023/golang/lab4"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := Lab4.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := Lab4.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
