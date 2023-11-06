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

y := (sinCubed + cosCubed) * lnX

return y
}
package main

import (
"fmt"
"math"
)

func calculateY(x float64) float64 {
return (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
}

func main() {
x1 := 0.2
x2 := 0.3
x3 := 0.38
x4 := 0.43
x5 := 0.57

y1 := calculateY(x1)
y2 := calculateY(x2)
y3 := calculateY(x3)
y4 := calculateY(x4)
y5 := calculateY(x5)

fmt.Printf("y1 = %.4f\n", y1)
fmt.Printf("y2 = %.4f\n", y2)
fmt.Printf("y3 = %.4f\n", y3)
fmt.Printf("y4 = %.4f\n", y4)
fmt.Printf("y5 = %.4f\n", y5)
}
