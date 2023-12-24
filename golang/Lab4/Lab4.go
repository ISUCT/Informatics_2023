package Lab4

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func Lab4() {
	fmt.Println("Мутовкин Илья Алексеевич")
	
	taskAResults := internal.ZadachkamathForRange()
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := internal.ZadachkamathForValues()
	fmt.Println("Результаты TaskB:", taskBResults)
}
