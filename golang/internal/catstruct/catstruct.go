package catstruct

import "fmt"

type cat struct {
	age   int
	name  string
	color string
}

func (c cat) GetAge() int {
	return c.age
}

func NewCat(setAge int, setName, setColor string) cat {
	return cat{
		age:   setAge,
		name:  setName,
		color: setColor,
	}
}

func (c *cat) SetAge(age int) error {
	if age >= 0 && age <= 15 {
		c.age = age
		return nil
	}
	return fmt.Errorf("invalid cat age")
}

func (c cat) GetColor() string {
	return c.color
}

func (c cat) GetName() string {
	return c.name
}

func (c *cat) SetName(name string) {
	c.name = name
}
