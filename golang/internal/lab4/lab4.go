package lab4

import "math"

func task_solution(a, b, x float64) float64 {
	result := 1 + (math.Pow(math.Log10(x/a), 2))/b - math.Pow(math.E, x/a)
	return result
}

func Task_A(x_beg, x_end, x_stp, a, b float64) []float64 {
	list_answers := make([]float64, 0)
	for x := x_beg; x <= x_end; x += x_stp {
		list_answers = append(list_answers, task_solution(a, b, x))
	}
	return list_answers
}

func Task_B(xList []float64, a, b float64) []float64 {
	list_answers := make([]float64, 0)
	for i := range xList {
		list_answers = append(list_answers, task_solution(a, b, xList[i]))
	}
	return list_answers
}