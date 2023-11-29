package structcat

import (
	"fmt"
)

type cat struct {
	age    int
	gender string
	breed  string
	name   string
}

func (c cat) GetAge() int {
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

func (c *cat) AssignAge(age int) error {
	if age > 0 && age <= 17 {
		c.age = age
		return nil
	}
	return fmt.Errorf("I'm afraid there's no cat")
}

func (c *cat) AssignName(name string) {
	c.name = name
}

func (c *cat) AssignBreed(breed string) {
	c.breed = breed
}

func NewCat(setAge int, setGender, setBreed, setName string) cat {
	return cat{
		age:    setAge,
		gender: setGender,
		breed:  setBreed,
		name:   setName,
	}
}
