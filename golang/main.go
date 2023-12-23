package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Кочетов данил Романович")

	fmt.Println("Lab 4")

	a := 0.05
	b := 0.06
	x_beg := 0.2
	x_end := 0.95
	x_stp := 0.15

	xList := []float64{0.08, 0.26, 0.35, 0.41, 0.53}

	fmt.Println(lab4.Task_A(x_beg, x_end, x_stp, a, b))
	fmt.Println(lab4.Task_B(xList, a, b))

	fmt.Println("Lab 5")

	table, err := lab5.NewTable(0, 1.34, 0.96)
	checkForError(err)

	fmt.Printf("Table length is %f\n", table.GetLength())
	fmt.Printf("Table width is %f\n", table.GetWidth())
	fmt.Printf("Table height is %f\n", table.GetHeight())
}
