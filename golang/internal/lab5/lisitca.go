package lab5

import (
	"errors"
	"fmt"
)

// структура лисица
type lisitca struct {
	name  string
	color colors
	age   int
}

// проверка и установка возраста
var ErrorAgeIsIncorrect = errors.New("возраст указан не верно")

func (l *lisitca) SetAge(age int) error {
	if age < 0 || age > 30 {
		return ErrorAgeIsIncorrect
	}
	l.age = age
	return nil
}

func (l *lisitca) GetAge() int {
	return l.age
}

// проверка и установка цвета
var ErrorColorIsIncorrect = errors.New("цвет указан не верно")

type colors string

const (
	orange colors = "оранжевый"
	white  colors = "белый"
	red    colors = "красный"
	gray   colors = "серый"
	brown  colors = "бурый"
)

func (l *lisitca) SetColor(color colors) error {
	switch color {
	case orange, white, red, gray, brown:
		l.color = color
		return nil
	}
	return ErrorColorIsIncorrect
}

func (l *lisitca) GetColor() int {
	return l.age
}

// имя
func (l *lisitca) SetName(name string) {
	l.name = name
}

func (l *lisitca) GetName() string {
	return l.name
}

// создание нового персонажа лисицы
func NewAnimalLisitca(name string, color colors, age int) (*lisitca, error) {
	animal := &lisitca{
		name: name,
	}
	if err := animal.SetAge(age); err != nil {
		return nil, err
	}
	if err := animal.SetColor(color); err != nil {
		return nil, err
	}
	return animal, nil
}

// ещё методы
func (l *lisitca) Voice(voice string) string {
	return fmt.Sprintf(
		"лисица %s говорит %s",
		l.name,
		voice,
	)
}

func (l *lisitca) Calling() string {
	return fmt.Sprintf(
		"вы подозвали лисицу %s, которая имеет %s цвет окраса",
		l.name,
		l.color,
	)
}
