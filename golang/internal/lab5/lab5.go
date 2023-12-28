package lab5

import (
	"fmt"
	"log"
)

type country struct {
	name       string
	square     int64
	population int64
}

func StructCreate(newName string, newSquare, newPopulation int64) country {
	s := country{
		name: newName,
	}
	err := s.setSquare(newSquare)
	if err != nil {
		log.Fatal(err)
	}
	err = s.setPopulation(newPopulation)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (c *country) setSquare(square int64) error {
	if square < 510072000 {
		c.square = square
		return nil
	} else {
		return fmt.Errorf("Старана не может быть площадью больше площади Земли")
	}
}

func (c *country) setPopulation(population int64) error {
	if population < 8028504258 {
		c.population = population
		return nil
	} else {
		return fmt.Errorf("Население страны не может быть больше населения всей Земли")
	}
}

func InfCountry(c *country) {
	fmt.Println("Название страны:", c.name)
	fmt.Println("Площадь страны:", c.square)
	fmt.Println("Население страны:", c.population)
}
