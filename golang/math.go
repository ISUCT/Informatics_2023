package main

import (
"fmt"
"math"
)

func main() {
a := 1.35
b := 0.98
x := 0.35

koren3 := math.Cbrt(a*x + b)
logorifm1 := math.Log10(x)
logorifm2 := logorifm1 * logorim1
resultat := koren3 / logorifm2
fmt.Println("Результат второй формулы - ", resultat)
}
