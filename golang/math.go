package main

import (
"fmt"
"math"
)

func main() {
startX := 0.11
endX := 0.36
step := 0.05

x := startX
for x <= endX {
result := calculateFunction(x)
fmt.Printf("x = %.2f, y = %.4f\n", x, result)
x += step
}
}

func calculateFunction(x float64) float64 {
sinCubed := math.Pow(math.Sin(x), 3)
cosCubed := math.Pow(math.Cos(x), 3)
lnX := math.Log(x)

return (sinCubed + cosCubed) * lnX
}
