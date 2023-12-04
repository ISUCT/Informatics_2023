package main

import (
	"fmt"
	"log"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type fox struct {
	age      int
	species  string
	isHungry bool
}

func newFox(age int, species string, isHungry bool) (fox, error) {
	var f fox = fox{
		isHungry: isHungry,
		species:  species,
	}
	var err = f.setAge(age)
	return f, err
}

func (f *fox) setAge(age int) error {
	if age >= 0 && age <= 15 {
		f.age = age
		return nil
	}
	return fmt.Errorf("fox has invalid age")
}

func (f fox) getAge() int {
	return f.age
}

func (f fox) getSpeices() string {
	return f.species
}

func (f fox) getView() error {
	asciiCode, err := os.ReadFile("Fox ASCII.txt")

	if err != nil {
		return err
	}

	fmt.Println(string(asciiCode))
	return nil
}

func (f *fox) feedTheFox() {
	if f.isHungry {
		f.isHungry = false
		fmt.Printf("%s is not hungry anymore", f.species)
	} else {
		fmt.Printf("%s doesn't want to eat", f.species)
	}
}

func main() {
	redFox, err := newFox(2, "Red Fox", false)
	errorCheck(err)
	fennecFox, err := newFox(3, "Fennec Fox", true)
	errorCheck(err)

	redFox.getView()
	fmt.Printf("This fox is %d years old\n", redFox.getAge())
	fmt.Printf("This fox's species is %s\n", redFox.getSpeices())

	fmt.Println("---")

	fmt.Printf("This fox's species is %s\n", fennecFox.getSpeices())
	fennecFox.feedTheFox()
}
