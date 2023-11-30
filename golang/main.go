package main

import (
  "fmt"
  "isuct.ru/informatics2022/internal/lab4"
)

func main() {
  fmt.Println("Васильева Вероника Дмитриевна")
  fmt.Println("Задача А")
  x,y := lab4.TaskA(1.25,3.25,0.4)
  for i := range y {
      fmt.Println("x=", x[i], "y=", y[i])
  }
  fmt.Println("Задача B")
  var list []float64 = []float64{1.84, 2.71, 3.81, 4.56, 5.62}
  x,y := lab4.TaskA(list)
  for i := range y{
      fmt.Println("x=", x[i], "y=", y[i])
  }
}
