package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab5"
)

// вариант 12
func main() {
	human_1, err := lab5.Newperson(37, "Сергей", "муж")
	lab5.CheckError(err)
	human_1.Set_Name("Константин")
	human_1.Set_Age(40)

	fmt.Printf("%s - %d\n", human_1.Get_Name(), human_1.Get_Age())
	human_2, err := lab5.Newperson(18, "Ульяна", "жен")
	lab5.CheckError(err)
	fmt.Printf("%s - %d\n", human_2.Get_Name(), human_2.Get_Age())
	human_2.Death_person()
}
