package lab5

import "fmt"

type cat struct {
	age   uint64
	breed string
	name  string
}

func CreateStruct(age uint64, breed, name string) cat {
	return cat{age, breed, name}
}

func SetAge(new_age uint64, c *cat) uint64 {
	c.age = new_age
	return c.age
}
func SetBreed(new_breed string, c *cat) string {
	c.breed = new_breed
	return c.breed
}
func SetName(new_name string, c *cat) string {
	c.name = new_name
	return c.name
}
func PrintInfCat(c *cat) {
	fmt.Println("Cat's name is", c.name, "\nCat's breed is", c.breed, "\nCat's age is", c.age, "\n")
}
