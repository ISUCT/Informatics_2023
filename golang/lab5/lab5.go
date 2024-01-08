package lab5

import (
	"errors"
	"fmt"
)

// структура лисичка
type Fox struct {
	name  string
	age   int
	color colors
}

func (f Fox) GetView() string {
	return `
	 /\   /\   
	//\\_//\\     ____
	\_     _/    /   /
	 / * * \    /^^^]
	 \_\O/_/    [   ]
	  /   \_    [   /
	  \     \_  /  /
	   [ [ /  \/ _/
	  _[ [ \  /_/ 
`
}
func (f Fox) GetName() string {
	return f.name
}
func (f Fox) GetAge() int {
	return f.age
}
func (f Fox) GetColor() colors {
	return f.color
}

// проверка возраста
func (f *Fox) SetAge(age int) error {
	if age >= 0 && age < 33 {
		f.age = age
		return nil
	}
	ErrorAgeIsIncorrect := errors.New("foxes don't live so long")
	return ErrorAgeIsIncorrect
}

// тип colors - установка цвета и проверка
type colors string

const (
	red   colors = "red"
	white colors = "white"
	grey  colors = "grey"
	black colors = "black"
)

// проверка цвета
func (f *Fox) SetColor(color colors) error {
	switch color {
	case red, white, grey, black:
		f.color = color
		return nil
	}
	ErrorColorIsIncorrect := errors.New("foxes don't have such color")
	return ErrorColorIsIncorrect
}
func NewFox(name string, age int, color colors) (Fox, error) {
	var f Fox = Fox{
		name:  name,
		age:   age,
		color: color}
	var err = f.SetAge(age)
	if err != nil {
		return f, err
	} else if err = f.SetColor(color); err != nil {
		return f, err
	}
	return f, err
}
func (f Fox) DisplayFox() {
	fmt.Printf("   Name - %s\n   Age - %d\n   Color - %s", f.name, f.age, f.color)
}
