package structs

import (
	"fmt"
	"strconv"
)

type phone struct {
	number   int
	provider string
	name     string
}

func (a phone) GetNumber() int {
	return a.number
}
func CreateStruct(number int, provider, name string) phone {
	return phone{number, provider, name}

}
func (a phone) GetName() string {
	return a.name
}
func (a phone) GetProvider() string {
	return a.provider
}
func (a *phone) SetNumber(number int) error {
	if len(strconv.Itoa(number)) != 11 {
		fmt.Printf("Невозможно установить номер - длина номера телефона должна быть равной 11. \n")
		return nil
	}
	a.number = number
	fmt.Printf("Пользователю %s добавлен номер. \n", a.GetName())
	return nil
}
func (a *phone) SetProvider(provider string) {
	a.provider = provider
	fmt.Printf("Пользователю %s установлен его провайдер. \n", a.GetName())
}
