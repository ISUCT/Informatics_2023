package personStruct

import (
	"errors"
	"fmt"
)
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
	var ErrWrongAge = errors.New("age is higher or lower than bound")
	return ErrWrongAge
}

func (p *person) SetGender(gender string) error {
	if gender == "M" || gender == "F" || gender == "Other" || gender == "Unknown" {
		p.gender = gender
		return nil
	}
	var ErrWrongGender = errors.New("gender can only be M/F/Other/Unknown")
	return ErrWrongGender
}

func NewPerson(name, gender string, age int) (person, error) {
	var p person = person {
		age: age,
		name: name,
		gender: gender,
	}
	var err = p.SetAge(age)
	if err != nil {
	    return p, err
	} else if err = p.SetGender(gender); err != nil {
		return p, err
	}
	return p, err
}

func (p person) KnowGender() {
	fmt.Printf("%s's gender - %s.\n", p.name, p.gender)
}

func (p person) GoToWalk(time int) (string, error) {
	if time <= 12 && time >= 0 {
		return ("At " + fmt.Sprint(time) + " o'clock " + p.name + " goes for a walk.\n"), nil
	}
	var ErrWrongTime = errors.New("time can't be more than 12 or less than 0")
	return " ", ErrWrongTime
}