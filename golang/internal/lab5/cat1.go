package main

import "fmt"

type cat struct {
	name string
	age  int
  colour string
}

func addname() int {
	return cat.name
}

func main() {
	var cat1 = cat{name: "Барсик", age: 3}

	// 1 кошачий год ~ 15 человеческим годам жизни

	vozrast := cat1.age * 15
	fmt.Println(cat1.name, cat1.age, vozrast)

	cat1.name = "Пушок"
	cat1.age = 5
	fmt.Println("Имя кота:", cat1.name, ".", "Возраст кота:", cat1.age, ",", "по переводу на человеческий:", vozrast)
	//Последний вариант, для более конкрентного обозначения переменных

}
