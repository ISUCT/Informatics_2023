package main

import (
	"fmt"
	"informatics2023/golang/internal"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := lab4.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := lab4.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
