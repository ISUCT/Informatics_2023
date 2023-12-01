package employeestruct

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
	date string
}

func Newemployee(setName string, setAge int, setDate string) Employee {
	return Employee{
		name: setName,
		age:  setAge,
		date: setDate,
	}
}

func (e Employee) Getname() string {
	return e.name
}

func (e Employee) Getage() int {
	return e.age
}

func (e Employee) Getdate() string {
	return e.date
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetAge(age int) error {
	if age > 17 && age < 80 {
		e.age = age
		return nil
	}
	return fmt.Errorf("age is incorrect")
}

func (e *Employee) SetDate(date string) {
	e.date = date
}
