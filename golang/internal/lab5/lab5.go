package lab5

import (
	"fmt"
)

type city struct {
	population int
	country string
}

func (c city) GetPopulation() int {
	return c.population
}

func NewCity(setPopulation int, setCountry string) city {
	return city{
		population: setPopulation,
		country: setCountry
	}
}

func (c *city) SetPopulation(population int) error {
	if population >= 0 && population <= 400 {
		c.population = population
		return nil
	}
	return fmt.Errorf("Invalid city population")
}

func (c city) GetCountry() string {
	return c.country
}

func (c city) GetTheCity() {
	fmt.Println("Welcome!")
}