package lab5

import (
	"fmt"
)

type Lisica struct {
	Name  string
	Color string
	Age   int
}

func (l *Lisica) SetName(name string) {
	l.Name = name
}

func (l *Lisica) SetColor(color string) error {
	if color != "orange" && color != "red" && color != "gray" && color != "brown" {
		return fmt.Errorf("указан неверный цвет")
	}
	l.Color = color
	return nil
}

func (l *Lisica) SetAge(age int) error {
	if age >= 0 {
		l.Age = age
		return nil
	}
	return fmt.Errorf("указан неверный взраст")
}

func (l *Lisica) GetName() string {
	return l.Name
}

func (l *Lisica) GetColor() string {
	return l.Color
}

func (l *Lisica) GetAge() int {
	return l.Age
}

func (l *Lisica) Calling() {
	fmt.Printf("%s лисицу зовут", l.GetName())
}
