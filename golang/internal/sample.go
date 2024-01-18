package internal

import (
	"fmt"
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
func func_res(x float64) float64 {
	return (4.1*math.Cbrt(x) - 2.7*log(5, x)) / math.Pow(math.Log10(x-1), 3)
}

func Task_A(x_start, x_end, x_step float64) []float64 {
	var capacity float64 = (x_end - x_start) / x_step
	var capacity_int int = int(capacity)
	answer_a := make([]float64, 0, capacity_int)
	for i := x_start; i <= x_end; i += x_step {
		answer_a = append(answer_a, func_res(i))
	}
	return answer_a
}

func Task_B(slice_of_x_values []float64) []float64 {
	answer_b := make([]float64, 0, len(slice_of_x_values))
	for _, i := range slice_of_x_values {
		answer_b = append(answer_b, func_res(i))
	}
	return answer_b
}

type City struct {
	Country    string
	Population uint
	NameCity   string
}

func (c *City) Output() {
	fmt.Println("Страна:", c.Country)
	fmt.Println("Население:", c.Population)
	fmt.Println("Город:", c.NameCity)
}

func (c *City) SetCountry(country string) {
	c.Country = country
}

func (c *City) GetCountry() string {
	return c.Country
}
