package lab5

import "fmt"

type fox struct {
	image  string
	name   string
	age    int
	gender string
}

func (f fox) GetView() string {
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

func (f fox) GetName() string {
	return f.name
}

func (f fox) GetAge() int {
	return f.age
}

func (f fox) GetGender() string {
	return f.gender
}

func (f *fox) SetAge(age int) error {
	if age >= 0 && age < 33 {
		f.age = age
		return nil
	}
	return fmt.Errorf("Foxes don't live that long")
}

func (f *fox) SetGender(gender string) error {
	if gender == "male" || gender == "female" {
		f.gender = gender
		return nil
	}
	return fmt.Errorf("Fox can only be male or female")
}

func NewFox(Name, Gender string, Age int) (fox, error, error) {
	var f fox = fox{
		name:   Name,
		age:    Age,
		gender: Gender,
	}
	var err_a = f.SetAge(Age)
	var err_g = f.SetGender(Gender)
	return f, err_a, err_g
}

func (f fox) DisplayFox() {
	fmt.Printf("       Name - %s\n       Age - %d\n       Gender - %s", f.name, f.age, f.gender)
}
