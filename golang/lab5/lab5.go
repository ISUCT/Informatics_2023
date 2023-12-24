package lab5

import "fmt"

type dog struct {
	name  string
	breed string
	age   int
}

func NewDog(enter_name, enter_breed string, enter_age int) (dog, error) {
	d := dog{
		name:  enter_name,
		breed: enter_breed,
	}
	err := d.SetAge(enter_age)
	return d, err
}

func (d dog) GetName() string {
	return d.name
}

func (d dog) GetBreed() string {
	return d.breed
}

func (d dog) GetAge() int {
	return d.age
}

func (d *dog) SetAge(age int) error {
	if age >= 0 && age < 16 {
		d.age = age
		return nil
	}
	return fmt.Errorf("%s; you entered an incorrect age for the dog; %s", d.name, d.breed)
}
