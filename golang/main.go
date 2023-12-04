package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Вераксо Марк Владимирович 1/280")
	fmt.Println("Задача А с циклом: ",internal.сycle(1.28, 3.28, 0.4))
	fmt.Println("Задача В со срезом: ",internal.Srez([]float64{1.1, 2.4, 3.6, 1.7, 3.9}))
}
