package lab5

import (
	"errors"
)

const (
	ErrInvalidWeight = "неверно указан вес, пожалуйста, укажите реальное значение"
	ErrInvalidHeight = "неверно указан рост, пожалуйста, укажите реальное значение"
	ErrInvalidAge    = "неверно указан возраст, пожалуйста, укажите реальное значение"
)

type Person struct {
	weight float64
	height float64
	age    int
	name   string
}

func NewPerson(weight, height float64, age int, name string) (*Person, error) {
	p := &Person{
		weight: weight,
		height: height,
		age:    age,
		name:   name,
	}
	if err := p.SetWeight(weight); err != nil {
		return nil, err
	}
	if err := p.SetHeight(height); err != nil {
		return nil, err
	}
	if err := p.SetAge(age); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Person) GetWeight() float64 {
	return p.weight
}

func (p *Person) SetWeight(weight float64) error {
	if weight < 0 || weight > 1000 {
		return errors.New(ErrInvalidWeight)
	}
	p.weight = weight
	return nil
}

func (p *Person) GetHeight() float64 {
	return p.height
}

func (p *Person) SetHeight(height float64) error {
	if height < 0 || height > 500 {
		return errors.New(ErrInvalidHeight)
	}
	p.height = height
	return nil
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) error {
	if age < 0 || age > 200 {
		return errors.New(ErrInvalidAge)
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
