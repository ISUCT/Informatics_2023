package main

import (
	"log"
	"fmt"
	"isuct.ru/informatics2022/internal"

)

func main() {
	answer1 := internal.Cycle(1.28, 3.28, 0.4)
	answer2 := internal.Srez(2.5, []float64{1.1, 2.4, 3.6, 1.7, 3.9})
	fmt.Println("Вераксо Марк Владимирович 1/280")
	fmt.Println("1 Вариант")
	fmt.Println(answer1)
	fmt.Println(answer2)

	firstcat, error := internal.NewCat("Pushok", 5, "Black")
	checkForError(error)
	secondcat, error := internal.NewCat("Aegis", 7, "Yellow")
	checkForError(error)
	thirdcat, error :=internal.NewCat("Simba", 10, "White")

	firstcat.AddAge(9)
	checkForError(error)

	secondcat.AddName("Aegis")
	checkForError(error)

	thirdcat.AddColour("Gray")
	checkForError(error)

	fmt.Println("Первому коту мы обновили возраст на значение:", firstcat.AddAge(9))
	fmt.Println("Второму коту мы добавили новое имя:", firstcat.AddName("Aegis"))
	fmt.Println("Третьему коту мы добавили новый цвет:", firstcat.AddColour("Gray"))
	
	fmt.Println(firstcat)
	fmt.Println(secondcat)
	fmt.Println(thirdcat)



//Первый пример: первому коту мы добавили новый возраст - 9
//Второй пример: второму коту мы добавили новое имя - Aegis
//Третий пример: третьему коту мы добавили новый цвет - Gray
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

