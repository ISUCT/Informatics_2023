package main

import ( 
	"math"
	"fmt"
)

func main() {
	fmt.Println("Хабибулина Алина Михайловна")
}
// Лабораторная 3 - Вариант 20
func func_res(){
a := 0.8
b := 0.4
x_begin := 1.23
x_end := 7.23
x_delta := 1.2
fmt. Println("Высчитывание задач")
	for x := x_begin; x <= x_end; x = x + x_delta {
		var summ1 float64 = x + a
		var summ2 float64 = x + b
		root1 := math.Pow(math.Pow(summ1,2),root_nth_degree(3))
		root2 := math.Pow(math.Abs(summ2),root_nth_degree(5))
		summ_chisl := root1 + root2
		summ3 := math.Pow(x,2)-math.Pow(a+b,2)
		root3 := math.Pow(summ3, root_nth_degree(9))
		resolution:= summ_chisl/root3
return resolution
}
func root_nth_degree(degree float64) float64{
	return 1 / degree
}
}