package lab5

import (
	"fmt"
	"log"
)

type rabbit struct {
	age    int
	name   string
	weight float64
}

func Newrabbit(age int, name string, weight float64) (rabbit, error) {
	var r rabbit = rabbit{
		name:   name,
		weight: weight,
	}
	var err = r.Setage(age)
	return r, err
}

func (r *rabbit) Getage() int {
	return r.age
}

func (r *rabbit) Setage(age int) error {
	if age >= 0 && age <= 10 {
		r.age = age
		return nil
	}
	return fmt.Errorf("Rabbit %s's age is incorrect", (r.Getname()))
}

func (r *rabbit) Getname() string {
	return r.name
}

func (r *rabbit) Setname(name string) {
	r.name = name
}

func (r *rabbit) Getweight() float64 {
	return r.weight
}

func (r *rabbit) Setweight(weight float64) {
	r.weight = weight
}

func Rabbitinfo(r *rabbit) {
	fmt.Println("rabbit's age is", r.Getage(), "\nrabbit's name is", r.Getname(), "\nrabbit's weight is", r.Getweight())
}

func Checkforeror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
