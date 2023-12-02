package lab4

import "math"

func solution_task(a, b, x float64) float64 {
	result := math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)/math.Acos(x*b) + math.Exp(-x/2)
	return result
}

func Task_A(x_beg, x_end, x_stp, a, b float64) []float64 {
	list_answers := make([]float64, 0, int(math.Ceil((x_end-x_beg)/x_stp+1)))
	for x := x_beg; x <= x_end; x += x_stp {
		list_answers = append(list_answers, solution_task(a, b, x))
	}
	return list_answers
}

func Task_B(xList []float64, a, b float64) []float64 {
	list_answers := make([]float64, 0, len(xList))
	for i := range xList {
		list_answers = append(list_answers, solution_task(a, b, xList[i]))
	}
	return list_answers
}
