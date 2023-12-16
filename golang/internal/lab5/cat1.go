package internal

import "fmt"

type cat struct {
	name string
	age  int
    colour string
}

//первый метод addage для получателя cat1

func (cat1 *cat) addage(age int) error{
	if 0 <= age && age <= 18{
		cat1.age = age
		return nil
	}
	return fmt.Errorf("Ошибка. Возраст кота не точный.")
}

//второй метод addname для получателя cat1
func (cat1 *cat) addname() string{
	return cat1.name
}

//третий метод addcolour для получателя cat1
func (cat1 *cat) addcolour() string{
	return cat1.name
}

func addcat(namecat string, agecat int, colourcat string) (cat,error) {
	cat1 := cat{name: namecat, colour:colourcat}
	var error = cat1.addage(agecat)
	return cat1, error
}
