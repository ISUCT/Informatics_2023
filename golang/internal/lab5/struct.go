package lab5

import (
	"fmt"
)

type dog struct {
	age    int
	name   string
	gender string
}

func (d dog) GetAge() int {
	return d.age
}

func (d dog) GetName() string {
	return d.name
}

func (d dog) GetGender() string {
	return d.gender
}

func TheDog(dog_age int, dog_name, dog_gender string) (dog, error) {
	var d dog = dog{
		name:   dog_name,
		gender: dog_gender,
	}
	var err = d.SetAge(dog_age)
	return d, err
}

func (d *dog) SetAge(age int) error {
	if age > 0 && age < 15 {
		d.age = age
		return nil
	}
	return fmt.Errorf("Sobaka %s s defektom", d.GetName())
}
