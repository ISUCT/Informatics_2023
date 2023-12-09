package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := lab4.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := lab4.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
