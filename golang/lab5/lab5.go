package lab5

import (
	"fmt"
	"log"
)

type Rabbit struct {
	age    uint
	weigth float64
	color  string
	gender string
}

func New_Struct(new_age uint, new_w float64, new_color, new_gender string) Rabbit {
	var new_rabbit = Rabbit{
		color: new_color,
	}
	err := new_rabbit.set_weigth(new_w)
	if err != nil {
		log.Fatal(err)
	}
	err = new_rabbit.set_age(new_age)
	if err != nil {
		log.Fatal(err)
	}
	err = new_rabbit.set_gender(new_gender)
	if err != nil {
		log.Fatal(err)
	}
	return new_rabbit
}
func (r *Rabbit) set_age(new_age uint) error {
	if new_age < 15 {
		r.age = new_age
		return nil
	} else {
		return fmt.Errorf("некорректно введён возраст кролика")
	}
}
func (r *Rabbit) set_weigth(new_w float64) error {
	if (new_w >= 0.4) && (new_w <= 2.0) {
		r.weigth = new_w
		return nil
	} else {
		return fmt.Errorf("некорректно введён вес кролика")
	}
}
func (r *Rabbit) set_gender(new_gender string) error {
	if new_gender == "муж" || new_gender == "жен" {
		r.gender = new_gender
		return nil
	} else {
		return fmt.Errorf("некорректно введён гендер кролика")
	}
}
func Output(r *Rabbit) {
	fmt.Printf("Кролик весит %1.2f килограмм, имеет возраст %d лет, %s пол и %s цвет\n", r.weigth, r.age, r.gender, r.color)
}
