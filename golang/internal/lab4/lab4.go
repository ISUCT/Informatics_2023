package lab4

import "math"

func solve_task(a, b, x float64) float64 {
	result := math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))
	return result
}

func Task_A(x_beg, x_end, x_stp, a, b float64) []float64 {
	list_answers := make([]float64, 0)
	for x := x_beg; x <= x_end; x += x_stp {
		list_answers = append(list_answers, solve_task(a, b, x))
	}
	return list_answers
}

func Task_B(xList []float64, a, b float64) []float64 {
	list_answers := make([]float64, 0)
	for i := range xList {
		list_answers = append(list_answers, solve_task(a, b, xList[i]))
	}
	return list_answers
}
