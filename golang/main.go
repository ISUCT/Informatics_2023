package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab3"
	"isuct.ru/informatics2022/personStruct"
)

func ErrorTest(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Кремнев Юрий Юрьевич")

	a := 1.6
	x_begin := 1.2
	x_end := 3.7
	x_delta := 0.5

	lab3.Solve(a,x_begin,x_end,x_delta)

	person1, err1, err2 := personStruct.NewPerson("John", "M", 20)
	ErrorTest(err1)
	ErrorTest(err2)

	person2, err1, err2 := personStruct.NewPerson("Jessica", "Unknown", 19)
	ErrorTest(err1)
	ErrorTest(err2)

	person1.SetAge(25)
	err1 = person1.SetAge(25)
	ErrorTest(err1)

	person2.SetGender("F")
	err1 = person2.SetGender("F")
	ErrorTest(err1)

	fmt.Printf("%s's age - %d.\n", person1.GetName(), person1.GetAge())
	person2.KnowGender()

	str, err1 := person2.GoToWalk(12)
	if err1 == nil {
		fmt.Print(str)
	}
	ErrorTest(err1)
}