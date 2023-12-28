package lab5

import (
	"fmt"
	"log"
)

type cat struct {
	age   uint64
	breed string
	name  string
}

func CreateStruct(age_to_var uint64, breed_to_var, name_to_var string) (cat, error) {
	var newcat = cat{
		name:  name_to_var,
		breed: breed_to_var,
	}
	var err = newcat.SetAge(age_to_var)
	return newcat, err
}
func (c *cat) GetAge() uint64 {
	return c.age
}
func (c *cat) GetBreed() string {
	return c.breed
}
func (c *cat) GetName() string {
	return c.name
}

func (c *cat) SetAge(age uint64) error {
	if age < 20 && age > 0 {
		c.age = age
		return nil
	} else {
		return fmt.Errorf("Возраст кошки '%s' введен некоректно", (c.GetName()))
	}
}
func (c *cat) SetBreed(new_breed string) {
	c.breed = new_breed
}
func (c *cat) SetName(new_name string) {
	c.name = new_name
}
func PrintInfCat(c *cat) {
	fmt.Println("Cat's name is", c.GetName(), "\nCat's breed is", c.GetBreed(), "\nCat's age is", c.GetAge(), "\n------------------------------")
}
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
