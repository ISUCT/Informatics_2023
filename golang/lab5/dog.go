package dog

import (
	"errors"
)

type Dog struct {
	Name string
	Age int
	Weight float64
}

func NewDog(dog_name string, dog_age int, dog_weight float64) (Dog, error) {
	var d Dog = Dog{Name: dog_name}
	err := d.SetAge(dog_age)
	if err != nil {
		return Dog{}, err
	}
	err = d.SetWeight(dog_weight)
	if err != nil {
		return Dog{}, err
	}
	return d, nil
}

func (d *Dog) SetAge(age int) error {
	if age < 0 {
		return errors.New("Возраст не может быть отрицательным")
	}
	d.Age = age
	return nil
}

func (d *Dog) SetWeight(weight float64) error {
	if weight <= 0 {
		return errors.New("Вес не может быть меньше или равен 0")
	}
	d.Weight = weight
	return nil
}

func (d Dog) GetAge() int {
	return d.Age
}

func (d Dog) GetWeight() float64 {
	return d.Weight
}
