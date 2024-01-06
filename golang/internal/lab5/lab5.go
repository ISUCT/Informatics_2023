package lab5

import (
	"errors"
)

const (
	ErrIncorrectSpeed  = "incorrect speed"
	ErrIncorrectWeight = "incorrect weight"
)

type Car struct {
	speed  int
	weight int
	name   string
}

func (c *Car) SetSpeed(speed int) error {
	if speed >= 0 && speed <= 500 {
		c.speed = speed
		return nil
	}
	return errors.New(ErrIncorrectSpeed)
}

func (c *Car) SetWeight(weight int) error {
	if weight >= 1000 && weight <= 2000 {
		c.weight = weight
		return nil
	}
	return errors.New(ErrIncorrectWeight)
}

func (c *Car) SetName(name string) {
	c.name = name
}

func (c Car) GetSpeed() int {
	return c.speed
}

func (c Car) GetWeight() int {
	return c.weight
}

func (c Car) GetName() string {
	return c.name
}

func NewCar(speed int, weight int, name string) (*Car, error) {
	tmp := &Car{
		speed:  speed,
		weight: weight,
		name:   name,
	}
	if err := tmp.SetSpeed(speed); err != nil {
		return nil, err
	}
	if err := tmp.SetWeight(weight); err != nil {
		return nil, err
	}
	return tmp, nil
}
