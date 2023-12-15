package lab5

import (
	"fmt"
)

type Years_old int
type Grams int

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
	return fmt.Errorf("Incorrect age of bunny.")
}

func (b *Bunny) SetWeight(weight Grams) error {
	if weight >= 62 && weight <= 6000 {
		b.weight = weight
		return nil
	}
	return fmt.Errorf("Incorrect weight of bunny.")
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
	err1 := tmp.SetAge(age)
	err2 := tmp.SetWeight(weight)
	if (err1 != nil && err2 != nil) {
		return nil, fmt.Errorf("Incorrect age and weight of bunny!")
	} else {
	if (err1 != nil) {
		return nil, err1
	}
	if (err2 != nil) {
		return nil, err2
	}
	return tmp, nil
	}
}