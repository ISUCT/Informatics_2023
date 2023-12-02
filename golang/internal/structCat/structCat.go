package structcat

import (
	
)

type cat struct {
	age    uint8
	gender string
	breed  string
	name   string
}

func (c cat) GetAge() uint8 {
	return c.age
}

func (c cat) GetGender() string {
	return c.gender
}

func (c cat) GetBreed() string {
	return c.breed
}

func (c cat) GetName() string {
	return c.name
}

func (c *cat) AssignName(name string) {
	c.name = name
}

func (c *cat) AssignBreed(breed string) {
	c.breed = breed
}

func (c *cat) AssignAge(age uint8) {
	c.age = age
}

func NewCat(setAge uint8, setGender, setBreed, setName string) cat {
	return cat{
		age:    setAge,
		gender: setGender,
		breed:  setBreed,
		name:   setName,
	}
}
