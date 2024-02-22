package laba5
import (
	"errors"
)
type (
	Year int
	Killogram int
)
const (
	ErrIncorrectYear = "неправильный возраст"
	ErrIncorrectKillogram = "неправильный вес"
)
type Dog struct {
	age Year 
	ves Killogram
	poroda string
}
func (d * Dog) SetAge(age Year) error{
	if age >= 0 && age <= 20{
		d.age = age
		return nil
	}
	return errors.New(ErrIncorrectYear)
}
func (d * Dog) SetVes(ves Killogram) error{
	if ves >= 5 && ves <= 65{
		d.ves = ves
		return nil
	}
	return errors.New(ErrIncorrectKillogram)
}
func (d * Dog) SetPoroda(poroda string){
	d.poroda = poroda
}
func (d Dog) GetAge() Year{
	return d.age
}
func (d Dog) GetVes() Killogram{
	return d.ves
}
func (d Dog) GetPoroda() string{
	return d.poroda
}
func NewDog(age Year, ves Killogram, poroda string)(*Dog, error){
	tmp := &Dog{
		age: age,
		ves : ves, 
		poroda: poroda,
	}
	if err := tmp.SetAge(age); err != nil{
		return nil, err
	}
	if err := tmp.SetVes(ves); err != nil{
		return nil, err
	}
	return tmp, nil
}