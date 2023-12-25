package lab5

import (
	"fmt"
	"log"
)

type person struct {
	age          int
	name, gender string
}

func Newperson(ageVar int, nameVar, gendefVar string) (person, error) {
	var p person = person{
		name:   nameVar,
		gender: gendefVar,
	}
	var err = p.Set_Age(ageVar)
	return p, err
}

func (p *person) Get_Age() int {
	return p.age
}

func (p *person) Get_Name() string {
	return p.name
}

func (p *person) Get_Gender() string {
	return p.gender
}

func (p *person) Set_Age(age int) error {
	if age >= 0 && age <= 130 {
		p.age = age
		return nil
	}
	return fmt.Errorf("Возраст человека '%s' введен некоректно", (p.Get_Name()))
}

func (p *person) Set_Name(name string) {
	p.name = name
}

func (p *person) Set_Gender(gender string) {
	p.gender = gender
}

func (p *person) Death_person() {
	if p.gender == "муж" {
		fmt.Printf("%s - был найден убитым\n", p.name)
	} else if p.gender == "жен" {
		fmt.Printf("%s - была найдена убитой\n", p.name)
	} else {
		fmt.Printf("%s - было найдено убитым\n", p.name)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
