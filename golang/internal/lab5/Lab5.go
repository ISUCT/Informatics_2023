package Lab5

import (
	"errors"
	"fmt"
)

type Mouse struct {
	Gender    string
	MouseName string
	Age       int
	Color     string
}

func (m *Mouse) GetGender() string {
	return m.Gender
}

func (m *Mouse) SetGender(Gender string) {
	m.Gender = Gender
}

func (m *Mouse) GetMouseName() string {
	return m.MouseName
}

func (m *Mouse) SetMouseName(MouseName string) {
	m.MouseName = MouseName
}

func (m *Mouse) GetAge() (int, error) {
	if m.Age < 0 {
		return 0, errors.New("мышь ещё не родилась, приходите позже")
	}
	if m.Age > 10 {
		return 0, errors.New("мышь умрёт от старости, лучше передумайте")
	}
	return m.Age, nil
}

func (m *Mouse) SetAge(Age int) {
	m.Age = Age
}

func (m *Mouse) GetColor() string {
	return m.Color
}

func (m *Mouse) SetColor(Color string) {
	m.Color = Color
}

func Lab5() *Mouse {
	fmt.Println("Выберите пол:")
	fmt.Println("1. Девочка")
	fmt.Println("2. Мальчик")
	var GenderChoice int
	fmt.Scanln(&GenderChoice)
	var Gender string
	switch GenderChoice {
	case 1:
		Gender = "Девочка"
	case 2:
		Gender = "Мальчик"
	default:
		fmt.Println("Я конечно всё понимаю...но боевой вертолёт")
	}
	MyMouse := Mouse{}
	MyMouse.SetGender(Gender)

	if Gender == "Девочка" {
		fmt.Println("Выберите вариант:")
		fmt.Println("1. Гаечка")
		fmt.Println("2. Сирень")
		fmt.Println("3. Марта")
		var nameChoice int
		fmt.Scanln(&nameChoice)

		switch nameChoice {
		case 1:
			MyMouse.SetMouseName("Гаечка")
		case 2:
			MyMouse.SetMouseName("Сирень")
		case 3:
			MyMouse.SetMouseName("Марта")
		default:
			fmt.Println("Имя не имя")
		}
	} else if Gender == "Мальчик" {
		fmt.Println("Выберите вариант:")
		fmt.Println("1. Болт")
		fmt.Println("2. Шуруп")
		fmt.Println("3. Гвоздь")
		var nameChoice int
		fmt.Scanln(&nameChoice)

		switch nameChoice {
		case 1:
			MyMouse.SetMouseName("Болт")
		case 2:
			MyMouse.SetMouseName("Шуруп")
		case 3:
			MyMouse.SetMouseName("Гвоздь")
		default:
			fmt.Println("Имя не имя")
		}
	}

	var AgeChoice int
	fmt.Println("Введите возраст мышки:")
	fmt.Scanln(&AgeChoice)
	MyMouse.SetAge(AgeChoice)
	fmt.Println("Какого цвета ваша мышка:")
	fmt.Println("1. Коричневая")
	fmt.Println("2. Чёрная")
	fmt.Println("3. Белая")
	var ColorChoice int
	fmt.Scanln(&ColorChoice)
	switch ColorChoice {
	case 1:
		MyMouse.SetColor("Розовая")
	case 2:
		MyMouse.SetColor("Чёрная")
	case 3:
		MyMouse.SetColor("Белая")
	default:
		fmt.Println("Подумай ещё раз")
	}

	return &MyMouse
}
