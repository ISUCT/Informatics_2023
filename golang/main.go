package main

import (
	"fmt"
	"math"
)

func taskA(xn float64, xk float64) {
	var deltax float64 = 0.4
	var answersy = []float64{}
	for x := xn; x <= xk; x += deltax {
		var y float64 = math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1/4) + math.Pow(math.Log10(math.Pow(x, 2)), 1/3)
		answersy = append(answersy, y)
	}
	fmt.Println(answersy)
}
func taskB(x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) {
	var znX = [5]float64{x1, x2, x3, x4, x5}
	var answersy = []float64{}
	for i := 0; i < len(znX); i++ {
		x := znX[i]
		var y float64 = math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1/4) + math.Pow(math.Log10(math.Pow(x, 2)), 1/3)
		answersy = append(answersy, y)
	}
	fmt.Println(answersy)

}

func main() {
	fmt.Println("Бусыгин Андрей Михайлович")
	taskA(1.25, 3.25)
	taskB(1.84, 2.71, 3.81, 4.56, 5.62)
}
