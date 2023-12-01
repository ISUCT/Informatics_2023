package lab5

import (
	"fmt"
)

type cat struct {
	age int
	name string
	breed string
}

func (c cat) GetAge() int {
	return c.age
}

func NewCat(setAge int, setName, setBreed string) cat {
	return cat{
		age: setAge,
		name: setName,
		breed: setBreed,
	}
}

func (c *cat) SetAge(age int) error {
	if age >= 0 && age <= 25 {
		c.age = age
		return nil
	}
	return fmt.Errorf("Invalid cat age")
}

func (c cat) GetBreed() string {
	return c.breed
}

func 