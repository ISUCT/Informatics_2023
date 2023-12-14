package main

import (
	"errors"
	"fmt"
)

type Cat struct {
	age    int
	weight float64
	height float64
	name   string
}

func NewCat(age int, weight float64, height float64, name string) (*Cat, error) {
	if age < 0 || age > 200 || weight < 0 || weight > 200 || height < 0 || height > 200 {
		return nil, errors.New("некорректные данные")
	}

	return &Cat{
		age:    age,
		weight: weight,
		height: height,
		name:   name,
	}, nil
}

func (c *Cat) GetAge() int {
	return c.age
}

func (c *Cat) SetAge(age int) error {
	if age < 0 || age > 200 {
		return errors.New("некорректные данные")
	}

	c.age = age
	return nil
}

func (c *Cat) GetWeight() float64 {
	return c.weight
}

func (c *Cat) SetWeight(weight float64) error {
	if weight < 0 || weight > 200 {
		return errors.New("некорректные данные")
	}

	c.weight = weight
	return nil
}

func (c *Cat) GetHeight() float64 {
	return c.height
}

func (c *Cat) SetHeight(height float64) error {
	if height < 0 || height > 200 {
		return errors.New("некорректные данные")
	}

	c.height = height
	return nil
}

func (c *Cat) GetName() string {
	return c.name
}

func (c *Cat) SetName(name string) error {
	c.name = name
	return nil
}

func main() {
	cat, err := NewCat(5, 3.5, 25.6, "Stesha")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Возраст: %d\n", cat.GetAge())
	fmt.Printf("Вес: %.2f\n", cat.GetWeight())
	fmt.Printf("Рост: %.3f\n", cat.GetHeight())
	fmt.Println("Имя:", cat.GetName())
}
