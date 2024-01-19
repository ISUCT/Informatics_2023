package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	fmt.Println("Киселев Ярослав Денисович")
	fmt.Println("---")

	a1 := lab4.Task1(0.05, 0.06, 0.2, 0.95, 0.15)
	fmt.Print(a1)

	a2 := lab4.Task2(0.05, 0.06, []float64{0.15, 0.26, 0.37, 0, 48, 0.56})
	fmt.Println(a2)

	table, err := lab5.NewTable(100, 130, 50)

	if err == nil {
		fmt.Printf("Table's height is %d\n", table.GetHeight())
		fmt.Printf("Table's length is %d\n", table.GetLength())
		fmt.Printf("Table's width is %d\n", table.GetWidth())
	} else {
		fmt.Println(err)
	}
}
