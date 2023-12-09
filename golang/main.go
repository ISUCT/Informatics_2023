package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab4"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//Лабораторная 2
	fmt.Println("Коротков Данила Максимович")
	//Лабораторная 3,4 (9 Вариант)
	a := 1.1
	b := 0.09
	x_begin := 1.2
	x_end := 2.2
	x_delta := 0.2
	lab4.Solve(a, b, x_begin, x_end, x_delta)
}
