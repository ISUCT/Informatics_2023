package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/employeestruct"
	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	n1 := lab4.Task1(4.1, 2.7, 1.2, 5.2, 0.7)
	fmt.Println(n1)
	n2 := lab4.Task2(4.1, 2.7, 1.9, 2.15, 2.34, 2.73, 3.16)
	fmt.Println(n2)

	employee := employeestruct.Newemployee("Иван Иванов", 30, "01/01/2023")

	employee.SetName("Павел Павлов")

	fmt.Printf("Employee's age is %d\n", employee.Getage())
	fmt.Printf("Eployee's name is %s\n", employee.Getname())
	fmt.Printf("Date is %s\n", employee.Getdate())

	fmt.Println("Полина Костина")
}
