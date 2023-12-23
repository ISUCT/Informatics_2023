package lab5

import "fmt"

type Employee struct {
	name string
	age  int
	date string
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

func (e *Employee) SetDate(date string) {
	e.date = date
}

func Newemployee(name string, age int, date string) (Employee, error) {
	var e Employee = Employee{
		name: name,
		date: date,
	}
	var err = e.SetAge(age)
	return e, err
}

func (e *Employee) SetAge(age int) error {
	if age > 17 && age < 80 {
		e.age = age
		return nil
	}
	return fmt.Errorf(e.Getname(), "age is incorrect", e.Getdate())
}
