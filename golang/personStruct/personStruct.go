package personStruct

import "fmt"

type person struct {
	name string
	age int
	gender string
}

func (p person) GetName() string {
	return p.name
}

func (p person) GetAge() int {
	return p.age
}

func (p person) GetGender() string {
	return p.gender
}

func (p *person) SetAge(age int) error {
	if age > 0 && age < 122 {
		p.age = age
		return nil
	}
	return fmt.Errorf("age is higher or lower than bound")
}

func (p *person) SetGender(gender string) error {
	if gender == "M" || gender == "F" || gender == "Other" || gender == "Unknown" {
		p.gender = gender
		return nil
	}
	return fmt.Errorf("gender can only be M/F/Other/Unknown")
}

func NewPerson(Name, Gender string, Age int) (person, error, error) {
	var p person = person {
		age: Age,
		name: Name,
		gender: Gender,
	}
	var err1 = p.SetAge(Age)
	var err2 = p.SetGender(Gender)
	return p, err1, err2 
}

func (p person) KnowGender() {
	fmt.Printf("%s's gender - %s.\n", p.name, p.gender)
}

func (p person) GoToWalk(time int) (string, error) {
	if time <= 12 && time >= 0 {
		return ("At " + fmt.Sprint(time) + " o'clock " + p.name + "goes for a walk.\n"), nil
	}
	return " ", fmt.Errorf("time can't be more than 12") 
}