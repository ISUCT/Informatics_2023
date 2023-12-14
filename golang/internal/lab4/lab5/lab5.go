package lab5

import (
	"errors"
)

type Person struct {
	weight float64
	height float64
	age    int
	name   string
}

func NewPerson(weight float64, height float64, age int, name string) (*Person, error) {
	if weight <= 0 || weight > 1000 || height <= 0 || height > 500 || age < 0 || age > 200 {
		return nil, errors.New("неверно указаны значения, пожалуйста, укажите реальные значения")
	}
	p := &Person{
		weight: weight,
		height: height,
		age:    age,
		name:   name,
	}
	return p, nil
}

func (p *Person) GetWeight() float64 {
	return p.weight
}

func (p *Person) SetWeight(weight float64) error {
	if weight < 0 || weight > 200 {
		return errors.New("неверно указан вес, пожалуйста, укажите реальное значение")
	}
	p.weight = weight
	return nil
}

func (p *Person) GetHeight() float64 {
	return p.height
}

func (p *Person) SetHeight(height float64) error {
	if height < 0 || height > 200 {
		return errors.New("неверно указан рост, пожалуйста, укажите реальное значение")
	}
	p.height = height
	return nil
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) error {
	if age < 0 || age > 200 {
		return errors.New("неверно указан возраст, пожалуйста, укажите реальное значение")
	}
	p.age = age
	return nil
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}
