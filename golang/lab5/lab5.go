package lab5

import (
	"errors"
	"fmt"
)

type fox struct {
	name   string
	age    int
	gender string
}

func (f fox) GetView() string {
	return `
	 /\   /\   
	//\\_//\\     ____
	\_     _/    /   /
	 / * * \    /^^^]
	 \_\O/_/    [   ]
	  /   \_    [   /
	  \     \_  /  /
	   [ [ /  \/ _/
	  _[ [ \  /_/ 
	  
`
}

func (f fox) GetName() string {
	return f.name
}

func (f fox) GetAge() int {
	return f.age
}

func (f fox) GetGender() string {
	return f.gender
}

func (f *fox) SetAge(age int) error {
	if age >= 0 && age < 33 {
		f.age = age
		return nil
	}
	ErrWrongAge := errors.New("Foxes don't live that long")
	return ErrWrongAge
}

func (f *fox) SetGender(gender string) error {
	if gender == "male" || gender == "female" {
		f.gender = gender
		return nil
	}
	ErrWrongGender := errors.New("Fox can only be male or female")
	return ErrWrongGender
}

func NewFox(name, gender string, age int) (fox, error) {
	var f fox = fox{
		name:   name,
		age:    age,
		gender: gender,
	}
	var err = f.SetAge(age)
	if err != nil {
		return f, err
	} else if err = f.SetGender(gender); err != nil {
		return f, err
	}
	return f, err
}

func (f fox) DisplayFox() {
	fmt.Printf("       Name - %s\n       Age - %d\n       Gender - %s", f.name, f.age, f.gender)
}
