package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.5
	for x := 1.28; x <= 3.28; x += 0.4 {
		verh := (1 + (math.Sin(b*b*b+x*x*x) * math.Sin(b*b*b+x*x*x)))
		niz := math.Cbrt(b*b*b + x*x*x)
		result := verh / niz

		fmt.Println("Результат первой формулы из таблицы - ", result)
	}
}
