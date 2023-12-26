package cat

import (
	"fmt"
)

type cat struct {
	name  string
	age   uint8
	color string
}

func NewCat(name_cat string, age_cat uint8, color string) (cat, error) {
	var c cat = cat{
		name: name_cat,
	}
	c.SetColor(color)
	var err = c.SetAge(age_cat)
	return c, err
}

func (c *cat) Meow() {
	fmt.Printf("%s мяукнул", c.GetName())
}

func (c *cat) SetName(name string) {
	c.name = name
}

func (c *cat) SetColor(color string) {
	c.color = color
}

func (c *cat) SetAge(age uint8) error {
	if age <= 50 {
		c.age = age
		return nil
	} else {
		return fmt.Errorf("Кошки не живут больше 50 лет!")
	}
}

func (c *cat) GetName() string {
	return c.name
}

func (c *cat) GetAge() uint8 {
	return c.age
}

func (c *cat) GetColor() string {
	return c.color
}
