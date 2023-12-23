package lab5

import (
	"errors"
)

type Years_old int
type Grams int

const (
	ErrIncorrectWeight = "incorrect weight"
	ErrIncorrectAge = "incorrect age"
)
type Bunny struct {
	age Years_old
	weight Grams
	name string
}

func (b *Bunny) SetAge(age Years_old) error {
	if age >= 0 && age <= 10 {
		b.age = age
		return nil
	}
	return errors.New(ErrIncorrectAge)
}

func (b *Bunny) SetWeight(weight Grams) error {
	if weight >= 62 && weight <= 6000 {
		b.weight = weight
		return nil
	}
	return errors.New(ErrIncorrectWeight)
}

func (b *Bunny) SetName(name string) {
	b.name = name
}

func (b Bunny) GetAge() Years_old {
	return b.age
}

func (b Bunny) GetWeight() Grams {
	return b.weight
}

func (b Bunny) GetName() string {
	return b.name
}

func NewBunny(age Years_old, weight Grams, name string) (*Bunny, error) {
	tmp := &Bunny{
		age: age,
		weight: weight,
		name: name,
	}
	if err := tmp.SetAge(age); err != nil {
		return nil, err
	}
	if err := tmp.SetWeight(weight); err != nil {
		return nil, err
	}
	return tmp, nil
}