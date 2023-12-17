package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab3"
	"isuct.ru/informatics2022/personStruct"
)

func main() {
	fmt.Println("Кремнев Юрий Юрьевич")

	a := 1.6
	x_begin := 1.2
	x_end := 3.7
	x_delta := 0.5

	lab3.Solve(a,x_begin,x_end,x_delta)

	person1, err := personStruct.NewPerson("John", "M", 20)
	if err != nil {
		log.Fatal(err)
	}
	
	person2, err := personStruct.NewPerson("Jessica", "Unknown", 19)
	if err != nil {
		log.Fatal(err)
	}

	err = person1.SetAge(25)
	if err != nil {
		log.Fatal(err)
	}

	err = person2.SetGender("F")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s's age - %d.\n", person1.GetName(), person1.GetAge())
	person2.KnowGender()

	str, err := person2.GoToWalk(12)
	if err == nil {
		fmt.Print(str)
	} else if err != nil {
		log.Fatal(err)
	}
}