package main

import (
	"fmt"
	"log"

	dishstruct "isuct.ru/informatics2022/internal/dishStruct"
	"isuct.ru/informatics2022/internal/lab4"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal("Card declined")
	}
}

func main() {
	fmt.Println("Kochetkov Kirill")

	res1 := lab4.Task1(7.2, 4.2, 1.81, 5.31, 0.7)
	fmt.Println(res1)

	res2 := lab4.Task2(7.2, 4.2, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(res2)

	dish := dishstruct.NewDish(5000, "soup", "Borsch")

	var err error = dish.SetPrice(4999)
	checkForError(err)

	dish.SetName("Sorrel")

	fmt.Printf("Your dish's price is %d\n", dish.GetPrice())
	fmt.Printf("Its type is %s\n", dish.GetVid())
	dish.TipTheWaiter()
}
