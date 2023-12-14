package internal

import (
	"fmt"
	"math"
)

const a = 0.8
const b = 0.4

func Summ(a, b float64) float64 {
	return a + b
}

func Cycle(x_begin, x_end, x_step float64) []float64 {
	var slice = make([]float64, 0, int64((x_end-x_begin)/x_step+1))
	for x := x_begin; x <= x_end; x += x_step {
		slice = append(slice, problem_1(a, b, x))
	}
	return slice
}

func Massiv(chek []float64) []float64 {
	var answer = make([]float64, 0, len(chek))
	for _, val := range chek {
		answer = append(answer, problem_1(a, b, val))
	}
	return answer
}

func problem_1(a, b, x float64) float64 {
	result := (math.Pow(math.Pow((x-a), 2), drob(3.0)) + math.Pow(math.Abs(x+b), drob(5.0))) / (math.Pow((math.Pow(x, 2) - math.Pow((a+b), 2)), drob(9.0)))
	return result
}

func drob(y float64) float64 {
	result_1 := 1 / y
	return result_1
}

type Piggy struct {
	name  string
	age   int64
	breed string
	sex   string
	color string
}

func Structure() {
	var PiggyBlin Piggy = Piggy{name: "Rose", age: 4, breed: "Mangalitsa", sex: "female", color: "brown"}
	PiggyBlin.print()
	PiggyBlin.setName("Teodora")
	PiggyBlin.setColor("Pink")
	PiggyBlin.setBreed("Hereford")
}

func (pig Piggy) print() {
	fmt.Println("Имя", pig.name)
	fmt.Println("Возраст", pig.age)
	fmt.Println("Порода", pig.breed)
	fmt.Println("Пол", pig.sex)
	fmt.Println("Цвет", pig.color)
}

func (pig *Piggy) setName(name1 string) {
	pig.name = name1
}

func (pig *Piggy) setColor(color1 string) {
	pig.color = color1
}

func (pig *Piggy) setBreed(breed1 string) {
	pig.breed = breed1
}
