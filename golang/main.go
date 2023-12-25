package main

import (
	"fmt"
	"log"
	"isuct.ru/informatics2022/sample" 
	"isuct.ru/informatics2022/ lab5" 
)
func checkForError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
func main() {
	
	a := 2.5
	b := 4.6
	x_n := 1.1
	x_k := 3.6
	x_z := 0.5

	fmt.Println("Значение функции №18 при данных задачи А =", sample.TaskA(x_n, x_k, x_z, a, b))
	fmt.Println("Значение функции №18 при данных задачи Б =", sample.TaskB([]float64{1.2, 1.28, 1.36, 1.46, 2.35}, a, b))
	fmt.Println("Lab 5")

	ship, er := lab5.TheShip(22.8, 53.3, 32)
	checkForError(er)

	fmt.Printf("The length of the ship is %f meters", ship.GetLength())
	fmt.Printf("The displacement of the ship is %f tons", ship.GetDisplacement())
	fmt.Printf("The max speed of the ship is %d knots", ship.GetMaxSpeed())
}


