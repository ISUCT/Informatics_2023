package lab5

import (
	"fmt"
)

type Mouse struct {
	Name string
	Age  int64
	Sex  string
}

// funcGet
func (m Mouse) GetName() string {
	return m.Name
}

func (m Mouse) GetAge() int64 {
	return m.Age
}

func (m Mouse) GetSex() string {
	return m.Sex
}

// funcSet
func (m *Mouse) SetName(name1 string) {
	m.Name = name1
}

func (m *Mouse) SetAge(age int64) {
	m.Age = age
}

func (m *Mouse) SetSex(sex string) {
	m.Sex = sex
}

func NewMouse(name, sex string, age int64) (Mouse, string) {
	var m Mouse = Mouse{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
	str := " "
	if m.Age < 0 {
		str = "Указанный возраст мышки явлется недопустимым"
		return m, str
	}
	return m, str
}

func (m Mouse) GetInfo() {
	fmt.Println("Name:", m.GetName())
	fmt.Println("Age:", m.GetAge())
	fmt.Println("Sex:", m.GetSex())
}

func (m Mouse) MickeyM() {
	m.SetAge(92)
	m.SetName("Mickey Mouse")
	m.SetSex("Male")

	m.GetInfo()

	fmt.Println(`                                  .--,       .--,
			 ( (  \.---./  ) )
			  '.__/o   o\__.'
			     {=  ^  =}
			    >  '._.'  <
		___________.""-------"".____________
		|  o                            O    \
		|                                    /
		|  .  O Hello! I'm Mikkey Mouse =) o \
		|                                    /         __
		|                                    \     _.-'  '.
		|______________o__________o__________/ .-~^        '~--'
			    ___)( )(___        '-.___.'
			   (((__) (__)))`)
}

func (m Mouse) MinnieM() {
	m.SetName("Minnie Mouse")
	m.SetAge(82)
	m.SetSex("Female")

	m.GetInfo()

	fmt.Println(`                                  .--,       .--,
			 ( (  \.---./  ) )
			  '.__/o   o\__.'
		             {=  ^  =}
			    >  '._.'  <
		___________.""-------"".____________
		|  o                            O    \
		|                                    /
		|  . O  Hello! I'm Minnie Mouse ;) o \
		|                                    /         __
		|                                    \     _.-'  '.
		|______________o__________o__________/ .-~^        '~--'
			    ___)( )(___        '-.___.'
			   (((__) (__)))`)
}
