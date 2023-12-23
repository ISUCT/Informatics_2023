package lab5

import (
	"errors"
)

const (
	ErrInvalidWeight = "неверно указан вес мыши, пожалуйста, укажите действительное значение"
	ErrInvalidHeight = "неверно указан рост мыши, пожалуйста, укажите действительное значение"
	ErrInvalidAge    = "неверно указан возраст мыши, пожалуйста, укажите действительное значение"
)

type Mouse struct {
	weight float64
	height float64
	age    int
	name   string
}

func NewMouse(weight, height float64, age int, name string) (*Mouse, error) {
	m := &Mouse{
		weight: weight,
		height: height,
		age:    age,
		name:   name,
	}
	if err := m.SetWeight(weight); err != nil {
		return nil, err
	}
	if err := m.SetHeight(height); err != nil {
		return nil, err
	}
	if err := m.SetAge(age); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Mouse) GetWeight() float64 {
	return m.weight
}

func (m *Mouse) SetWeight(weight float64) error {
	if weight < 0 || weight > 30 {
		return errors.New(ErrInvalidWeight)
	}
	m.weight = weight
	return nil
}

func (m *Mouse) GetHeight() float64 {
	return m.height
}

func (m *Mouse) SetHeight(height float64) error {
	if height < 0 || height > 30 {
		return errors.New(ErrInvalidHeight)
	}
	m.height = height
	return nil
}

func (m *Mouse) GetAge() int {
	return m.age
}

func (m *Mouse) SetAge(age int) error {
	if age < 0 || age > 50 {
		return errors.New(ErrInvalidAge)
	}
	m.age = age
	return nil
}

func (m *Mouse) GetName() string {
	return m.name
}

func (m *Mouse) SetName(name string) {
	m.name = name
}
