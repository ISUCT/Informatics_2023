package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/laba4"
	"isuct.ru/informatics2022/internal/laba5"
)

func main() {
	fmt.Println("Яковлева Алёна")
	dog, err := laba5.NewDog(20, 65, "Овчарка")
	if err == nil {
		fmt.Printf("Dog age is %d\n", dog.GetAge())
		fmt.Printf("Dog ves is %d\n", dog.GetVes())
		fmt.Printf("Dog poroda is %s\n", dog.GetPoroda())
	}else{
		fmt.Println(err)
	}
	y1 := laba4.Task1(0.22, 0.92, 0.14)
	y2 := laba4.Task2([]float64{0.1, 0.35, 0.4, 0.55, 0.6})
	fmt.Println(y1)
	fmt.Println(y2)
}
