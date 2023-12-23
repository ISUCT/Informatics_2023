package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/internal/Lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func main() {
	//Work1
	fmt.Println("AnswerOnWork1:")
	resultWork1 := Lab4.Work1(2.0, 1.1, 0.08, 1.08, 0.2)
	fmt.Println(resultWork1)
	//Work2
	SliceWork2 := []float64{0.1, 0.3, 0.4, 0.45, 0.65}
	fmt.Println("AnswerOnWork2:")
	resultSlice := Lab4.Work2(SliceWork2, 2.0, 0)
	fmt.Println(resultSlice)
	}
	
	fmt.Println("Мазуров Даниил Алексеевич")

	ace, err := lab5.NewACE(1000, 2000, 3000, "ACE 123")
	if err != nil {
		fmt.Println("Ошибка создания ACE:", err)
		return
	}

	fmt.Println("Текущие значения ACE:")
	fmt.Printf("Firepower: %.2f\n", ace.GetFirepower())
	fmt.Printf("Speed: %.2f\n", ace.GetSpeed())
	fmt.Printf("Armor: %.2f\n", ace.GetArmor())
	fmt.Printf("Name: %s\n", ace.GetName())

	// Изменение значений через методы Set
	err = ace.SetFirepower(1500)
	if err != nil {
		fmt.Println("Ошибка при установке огневой мощи:", err)
		return
	}

	err = ace.SetSpeed(2500)
	if err != nil {
		fmt.Println("Ошибка при установке скорости:", err)
		return
	}

	err = ace.SetArmor(3500)
	if err != nil {
		fmt.Println("Ошибка при установке бронирования:", err)
		return
	}

	ace.SetName("ACE 456")
	fmt.Println("\nИзмененные значения ACE:")
	fmt.Printf("Firepower: %.2f\n", ace.GetFirepower())
	fmt.Printf("Speed: %.2f\n", ace.GetSpeed())
	fmt.Printf("Armor: %.2f\n", ace.GetArmor())
	fmt.Printf("Name: %s\n", ace.GetName())
}
