package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab5"
	// "isuct.ru/informatics2022/internal/lab4"
)

func main() {
	// lab4
	// fmt.Println("Васильева Вероника Дмитриевнa")
	// fmt.Println("Задача А")
	// x, y := lab4.TaskA(1.25, 3.25, 0.4)
	// for i := range y {
	// 	fmt.Println("x=", x[i], "y=", y[i])
	// }
	// fmt.Println("Задача B")
	// var list []float64 = []float64{1.84, 2.71, 3.81, 4.56, 5.62}
	// yB := lab4.TaskB(list)
	// for j := range yB {
	// 	fmt.Println("x=", list[j], "y=", yB[j])
	// }
	// lab5
	car1 := lab5.Mycar("BMW e34", 220, 1865000, 432057)
	car2 := lab5.Mycar("Lexus ls400", 130, 900000, 653285)
	fmt.Printf("Модель автомобиля: %s\n", car1.Getmodel())
	fmt.Printf("Максимальная скорость: %d\n", car1.Getmaxspeed())
	fmt.Printf("Цена %d\n", car1.Getprice())
	fmt.Printf("Пробег %d\n", car1.Getmileage())
	car2.Setmaxspeed(300)
	car2.Setmaxspeed(50000)
	car2.Setmileage(389000)
	fmt.Printf("Максимальная скорость: %d\n", car2.Getmaxspeed())
	fmt.Printf("Пробег %d\n", car2.Getmileage())
}
