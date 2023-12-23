package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
  fmt.Println("Жуков Даниил Евгеньевич")
	fmt.Println(internal.TaskA(0.7, 2.2, 0.3))
	fmt.Println(internal.TaskB([]float64{0.25, 0.36, 0.56, 0.94, 1.28}))
}
