package lab5

import "errors"

const (
	ErrAge    = "неверно указан возраст, пожалуйста, укажите реальное значение"
	ErrWeight = "неверно указан вес, пожалуйста, укажите реальное значение"
	ErrHeight = "неверно указан рост, пожалуйста, укажите реальное значение"
)

type Cat struct {
	age    int
	weight float64
	height float64
	name   string
}

func NewCat(age int, weight, height float64, name string) (*Cat, error) {
	c := &Cat{
		age:    age,
		weight: weight,
		height: height,
		name:   name,
	}
	if error := c.SetWeight(weight); error != nil {
		return nil, error
	}
	if error := c.SetHeight(height); error != nil {
		return nil, error
	}
	if error := c.SetAge(age); error != nil {
		return nil, error
	}
	return c, nil
}

func (c *Cat) GetAge() int {
	return c.age
}

func (c *Cat) SetAge(age int) error {
	if age < 0 || age > 200 {
		return errors.New(ErrAge)
	}

	c.age = age
	return nil
}

func (c *Cat) GetWeight() float64 {
	return c.weight
}

func (c *Cat) SetWeight(weight float64) error {
	if weight < 0 || weight > 200 {
		return errors.New(ErrWeight)
	}

	c.weight = weight
	return nil
}

func (c *Cat) GetHeight() float64 {
	return c.height
}

func (c *Cat) SetHeight(height float64) error {
	if height < 0 || height > 200 {
		return errors.New(ErrHeight)
	}

	c.height = height
	return nil
}

func (c *Cat) GetName() string {
	return c.name
}

func (c *Cat) SetName(name string) {
	c.name = name
}
