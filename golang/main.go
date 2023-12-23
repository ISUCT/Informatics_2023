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
	n1 := lab4.Task1(4.1, 2.7, 1.2, 5.2, 0.7)
	fmt.Println(n1)
	n2 := lab4.Task2(4.1, 2.7, []float64{1.9, 2.15, 2.34, 2.73, 3.16})
	fmt.Println(n2)

	fmt.Println("Полина Костина")

	employee, err := lab5.Newemployee("Павел Павлов", 30, "01/01/2023")
	checkForError(err)

	fmt.Printf("Employee's age is %d\n", employee.Getage())
	fmt.Printf("Eployee's name is %s\n", employee.Getname())
	fmt.Printf("Date is %s\n", employee.Getdate())
}
