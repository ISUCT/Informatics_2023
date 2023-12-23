package main

import (
	"fmt"

	lab4test "isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab4/lab5"
)

func main() {
	fmt.Println("Жуков Михаил Андреевич")
	fmt.Println("17 лет")
	//task1
	fmt.Println("AnswerOnTask1:")
	resultSliceTaskA := lab4test.Task1Slice(1.1, 0.09, 1.2, 2.2, 0.2)
	fmt.Println(resultSliceTaskA)
	//task2
	mySliceForTaskB := []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	fmt.Println("AnswerOnTask2:")
	resultSlice := lab4test.TaskB(mySliceForTaskB, 1.1, 0.09)
	fmt.Println(resultSlice)

	//lab5

	person, err := lab5.NewPerson(50, 190, 100, "Mr. Positive")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Привет, друг! Вот характеристика человека:")
	fmt.Printf("Вес: %.1f\n", person.GetWeight())
	fmt.Printf("Рост: %.1f\n", person.GetHeight())
	fmt.Printf("Возраст: %d\n", person.GetAge())
	fmt.Printf("Имя: %s\n", person.GetName())

	err = person.SetWeight(75)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = person.SetHeight(175)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = person.SetAge(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	person.SetName("John")
	fmt.Println("Обновленные значения:")
	fmt.Println("Привет, друг! Вот характеристика человека:")
	fmt.Printf("Вес: %.1f\n", person.GetWeight())
	fmt.Printf("Рост: %.1f\n", person.GetHeight())
	fmt.Printf("Возраст: %d\n", person.GetAge())
	fmt.Printf("Имя: %s\n", person.GetName())
}
