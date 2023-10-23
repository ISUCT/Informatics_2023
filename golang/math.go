package main

import (
	"fmt"
	"math"
)
func main() {
	b := 2.5
	for x := 1.28; x <= 3.28; x += 0.4 {

	}

	skoba := (math.pow(b, 3) + math.pow(x, 3))
	sinus := math.Sin(skoba)
	verh := (1 + sinus*sinus)
	niz := math.Cbrt(skoba)
	itog := verh / niz

	fmt.Println("Результат первой формулы из таблицы - ", itog)
}
