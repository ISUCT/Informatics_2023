package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Иванова Анастасия Евгеньевна")
	fmt.Println("Вариант 7")
	fmt.Println(internal.Task_A(3.2, 6.2, 0.6))
	fmt.Println(internal.Task_B([]float64{4.48, 3.56, 2.78, 5.28, 3.21}))
}
