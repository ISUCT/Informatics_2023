package structcat

import "fmt"

type cat struct {
	age    uint8
	gender string
	breed  string
	name   string
}

func NewCat(age uint8, gender, breed, name string) (*cat, error) {
	var c cat = cat{
		name:   name,
		gender: gender,
		breed:  breed,
	}

	check := c.AssignedAge(age)

	if check != nil {
		return nil, check
	}
	return &c, check
}

func (c *cat) AssignedAge(age uint8) error {
	if age >= 17 {
		return fmt.Errorf("I'm afraid there is't cat here")
	}
	c.age = age
	return nil
}

func (c *cat) AssignedGender(gender string) {
	c.gender = gender
}

func (c *cat) AssigneName(name string) {
	c.name = name
}

func (c *cat) AssignedBreed(breed string) {
	c.breed = breed
}

func (c cat) GetAge() uint8 {
	return c.age
}

func (c cat) GetBreed() string {
	return c.breed
}

func (c cat) GetName() string {
	return c.name
}

func (c cat) GetGender() string {
	return c.name
}
