package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/internal/LAB4"
	structComp "isuct.ru/informatics2022/internal/structureComp"
)

// Лабораторная 4 вариант 5
func main() {
	fmt.Println("Гвоздарёв Ярослав Владимирович")
	fmt.Println("задача A")
	xNach := 3.5
	xKon := 6.5
	shag := 0.6
	a := -2.5
	b := 3.4
	x, y := lab4.TaskA(xNach, xKon, shag, a, b)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("задача B")
	xArr := []float64{2.89, 3.54, 5.21, 6.28, 3.48}
	x, y = lab4.TaskB(xArr, a, b)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	//Лабораторная 5 вариант 5
	fmt.Println("Лабораторная 5 вариант 5")
	disk := structComp.Computer(1024, "Kingston", "SSD")
	fmt.Println(disk)
	disk.Changename("Gygabyte")
	disk.Changetype("HDD")
	fmt.Printf("Your disk volume is %d GB\n", disk.Getvolume())
	fmt.Print(disk)
}
