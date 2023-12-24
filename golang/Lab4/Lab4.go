package Lab4

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func Lab4() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	a := 4.1
	b := 2.7
	xn := 1.2
	xk := 5.2
	deltaX := 0.8
	values := []float64{1.9, 2.15, 2.34, 2.73, 3.16}

	taskAResults := internal.ZadachkamathForRange(a, b, xn, xk, deltaX)
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := internal.ZadachkamathForValues(a, b, values)
	fmt.Println("Результаты TaskB:", taskBResults)
}
