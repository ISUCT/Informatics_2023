package main

import (
	"fmt"
	"informatics2023/golang/internal/Lab"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := Lab.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := Lab.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
