package main

import (
  "fmt"
  "math"
)

func formula(x float64) float64 {
  a :=2.0 
  return ((math.Pow((math.Log10(a+x)),2)) / (math.Pow((a+x),2)))
}

func main() {
  fmt.Println("Фёдорова Ольга Андреевна")
  fmt.Println("Задача А")
  for i := 1.2; i <= 4.2; i += 0.6 {
    fmt.Println("x=", i, "y=", formula(i))
  }
  fmt.Println("Задача B")
  var list [5]float64 = [5]float64{1.16, 1.32, 1.47, 1.65, 1.93}
  for i := 0; i <= 4; i++ {
    fmt.Println("x=", list[i], "y=", formula(list[i]))
  }
}
