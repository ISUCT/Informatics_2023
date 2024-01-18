package lab5

import (
	"fmt"
	"errors"
)

type city struct {
	population int
	country    string
}

func (c city) GetPopulation() int {
	return c.population
}

func NewCity(setPopulation int, setCountry string) (*city, error) {
	c := &city{
		population: setPopulation,
		country:    setCountry,
	}

	if err := c.SetPopulation(setPopulation); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *city) SetCountry(country string) {
	c.country = country
}

func (c *city) SetPopulation(population int) error {
	if population >= 0 && population <= 400 {
		c.population = population
		return nil
	}
	return errors.New("invalid city population")
}

func (c city) GetCountry() string {
	return c.country
}

func (c *city) GetTheCity() {
	fmt.Println("Welcome!")
}
