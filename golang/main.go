package main

import (
	"fmt"

	dishstruct "isuct.ru/informatics2022/internal/dishStruct"
	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Kochetkov Kirill")

	res1 := lab4.Task1(7.2, 4.2, 1.81, 5.31, 0.7)
	fmt.Println(res1)

	res2 := lab4.Task2(7.2, 4.2, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(res2)

	//lab5
	dish := dishstruct.NewDish(5000, "soup", "Borsch")
	err := dish.SetPrice(5000)
	if err != nil {
		fmt.Println(err)
		return
	}

	dish.SetName("Sorrel")

	fmt.Printf("Your dish's price is %d\n", dish.GetPrice())
	fmt.Printf("Its type is %s\n", dish.GetVid())
	fmt.Printf("Its called is %s\n", dish.GetName())
	dish.TipTheWaiter()
}
