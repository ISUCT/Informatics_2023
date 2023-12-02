package main

// Задача под А
import (
	"fmt"
	"math"
)

func main() {
	a1 := 2.0
	b1 := 3.0
	n := 0.11
	k := 0.36
	d := 0.05

	for x1 := n; x1 <= k; x1 += d {
		y1 := math.Asin(math.Pow(x1, a1)) + math.Acos(math.Pow(x1, b1))
		fmt.Printf("x1 = %.2f, y1 = %.2f\n", x1, y1)
	}

	// Задача под Б

	a2 := 2.0
	b2 := 3.0
	x2 := []float64{0.08, 0.26, 0.35, 0.41, 0.53}
	y2 := make([]float64, 5)

	for i := 0; i < len(x2); i++ {
		y2[i] = math.Asin(math.Pow(x2[i], a2)) + math.Acos(math.Pow(x2[i], b2))
	}

	for i := 0; i < len(x2); i++ {
		fmt.Printf("x2%d = %.2f, y2%d = %.4f\n", i+1, x2[i], i+1, y2[i])
	}
}
