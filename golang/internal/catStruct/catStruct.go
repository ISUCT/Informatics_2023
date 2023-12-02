package catStruct

import (
	"fmt"
)

type cat struct {
	age   int
	name  string
	breed string
}

func (c cat) GetAge() int {
	return c.age
}

func NewCat(ageVar int, nameVar, breedVar string) (cat, error) {
	var c cat = cat{
		name:  nameVar,
		breed: breedVar,
	}
	var err = c.SetAge(ageVar)
	return c, err
}

func (c *cat) SetAge(age int) error {
	if age >= 0 && age <= 25 {
		c.age = age
		return nil
	}
	return fmt.Errorf("cat named \"%s\" has invalid age", c.GetName())
}

func (c cat) GetBreed() string {
	return c.breed
}

func (c cat) GetName() string {
	return c.name
}

func (c *cat) SetName(name string) {
	c.name = name
}

func (c cat) PetTheCat() {
	fmt.Printf("Cat %s says meow\n", c.GetName())
}
