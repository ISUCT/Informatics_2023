package main

import (
	"fmt"
	"math"
)

func assignment_1() {
	var a float64 = 0.8
	var b float64 = 0.4
	var res_b = []float64{}
	for i := 1.23; i < 7.23; i += 1.2 {
		x := i
		var y float64 = (math.Pow((x-a), (2.00/3.00)) + math.Pow(math.Abs(x+b), (1.00/5.00))) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), (1.00/9.00))
		res_b = append(res_b, y)
	}
	fmt.Println(res_b)
}
func assignment_2() {
	var a float64 = 0.8
	var b float64 = 0.4
	var xs = []float64{1.88, 2.26, 3.84, 4.55, -6.21}
	var res_b = []float64{}
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		var y float64 = (math.Pow((x-a), (2.00/3.00)) + math.Pow(math.Abs(x+b), (1.00/5.00))) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), (1.00/9.00))
		res_b = append(res_b, y)
	}
	fmt.Println(res_b)
}
func main() {
	assignment_1()
	assignment_2()
	fmt.Println("Серунин Денис Витальевич")
}
