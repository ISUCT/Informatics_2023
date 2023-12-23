package main

import (
	"fmt"
)

type Mouse struct {
	name string
	age  int64
	sex  string
}

// funcGet
func (a Mouse) getName() string {
	return a.name
}

func (a Mouse) getAge() int64 {
	return a.age
}

func (a Mouse) getSex() string {
	return a.sex
}

// funcSet
func (a *Mouse) setName(name1 string) {
	a.name = name1
}

func (a *Mouse) setAge(age int64) {
	a.age = age
}

func (a *Mouse) setSex(sex string) {
	a.sex = sex
}

func main() {
	mouse1 := Mouse{"Mickey Mouse", 92, "Male"}

	var AgeCheck bool

	if mouse1.age < 0 {
		AgeCheck = false
	} else {
		AgeCheck = true
	}

	if AgeCheck == true {

		fmt.Println("Name:", mouse1.getName())
		fmt.Println("Age:", mouse1.getAge())
		fmt.Println("Sex:", mouse1.getSex())

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
				   (((__) (__))) `)

	} else {
		fmt.Println(`Указанный возраст является недопустимым`)
	}

	fmt.Println("                                                                    ")
	fmt.Println("                                                                    ")

	mouse1.setName("Minnie Mouse")
	mouse1.setAge(82)
	mouse1.setSex("Female")

	if mouse1.age < 0 {
		AgeCheck = false
	} else {
		AgeCheck = true
	}

	if AgeCheck == true {
		fmt.Println("Name:", mouse1.getName())
		fmt.Println("Age:", mouse1.getAge())
		fmt.Println("Sex:", mouse1.getSex())

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
				   (((__) (__))) `)

	} else {
		fmt.Println("Указанный возраст является недопустимым")
	}
}

// func formula(a, x float64) float64 {
// 	return ((math.Pow((math.Log(a + x)), 2)) / (math.Pow((a + x), 2)))
// }

// func main () {
// fmt.Println("Солдатов Иван Сергеевич")
// fmt.Println("------------------------First_Exercise-------------------------------")
// a := 2.0
// x_end := 4.2
// x_step := 0.6
// x_start := 1.2
// for x := x_start; x <= x_end; x += x_step {
// 	fmt.Println(formula(a, x))
// }
// fmt.Println("------------------------Second_Exercise-------------------------------")
// array := []float64{1.16, 1.32, 1.47, 1.65, 1.93}
// for _, x := range array {
// 	fmt.Println(formula(a, x))
// }
// fmt.Println("------------------------Lab_4-------------------------------")
// lab4.Solve(a, x_start, x_end, x_step)
// }
