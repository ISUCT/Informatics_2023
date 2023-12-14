package main

import (
	"fmt"
	"isuct.ru/informatics2022/golang/"
)

func main() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := Lab4.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := Lab4.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
