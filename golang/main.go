package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := lab4.zadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := lab4.zadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
