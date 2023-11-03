package main

import (
  "fmt"
  "math"
)

func f(x float64) float64 {
  return (math.Sqrt(math.Sqrt(math.Abs(math.Pow(x, 2)-2.5))) + math.Cbrt(math.Log10(math.Pow(x, 2))))
}

func main() {
  fmt.Println("Васильева Вероника Дмитриевна")
  fmt.Println("Задача А")
  for i := 1.25; i <= 3.25; i += 0.4 {
    fmt.Println("x=", i, "y=", f(i))
  }
  fmt.Println("Задача B")
  var list [5]float64 = [5]float64{1.84, 2.71, 3.81, 4.56, 5.62}
  for i := 0; i <= 4; i++ {
    fmt.Println("x=", list[i], "y=", f(list[i]))
  }
}
