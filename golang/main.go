package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	fmt.Println("Жуков Даниил Евгеньевич")
	fmt.Println(internal.TaskA(0.7, 2.2, 0.3))
	fmt.Println(internal.TaskB([]float64{0.25, 0.36, 0.56, 0.94, 1.28}))
	rabbit1, err := lab5.Newrabbit(3, "Benjamin", 5)
	lab5.Checkforeror(err)
	lab5.Rabbitinfo(&rabbit1)
	rabbit2, err := lab5.Newrabbit(7, "Dima", 10)
	lab5.Checkforeror(err)
	lab5.Rabbitinfo(&rabbit2)
	rabbit3, err := lab5.Newrabbit(20, "Alice", 15)
	lab5.Checkforeror(err)
	lab5.Rabbitinfo(&rabbit3)
}
