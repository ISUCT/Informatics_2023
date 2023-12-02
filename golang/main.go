package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/worker"
)

func main() {
	fmt.Println("Коротков Михаил Александрович | Вариант №13")
	fmt.Println(internal.Task_A(1.2, 5.2, 0.8))
	fmt.Println(internal.Task_B([]float64{1.9, 2.15, 2.34, 2.73, 3.16}))

	fmt.Println("Рабочий 1")
	work1, err := worker.NewWorker("Влад", "Хирург", 40000)
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Println(work1.GetSalary())
	work1.SetSalary(30000)
	fmt.Println(work1.GetSalary())
	work1.QuitJob()

	fmt.Println("")
	fmt.Println("Рабочий 2")
	work2, err := worker.NewWorker("Леонардо", "Уборщик", 13000)
	if (err != nil) {
		log.Fatal(err)
	}
	// Это не выводится, т.к. ошибка
	fmt.Println(work2.GetName())
}