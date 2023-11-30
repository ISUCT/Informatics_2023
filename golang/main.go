package main

import (
  "fmt"
  
  "isuct.ru/informatics2022/internal/lab4"
)

func main(){
  fmt.Println("Фёдорова Ольга Андреевна")
  fmt.Println("Задача A")
  x,y:=lab4.TaskA(1.2, 4.2, 0,6)
  for i := range y{
    fmt.Println("При x=", x[i],"y=",y[i])
  }
  fmt.Println("Задача B")
  var list []float64=[]float64{1.16, 1.32, 1.47, 1.65, 1.93}
  x,y=lab4.TaskB(list)
  for i := range y{
    fmt.Println("При x=",x[i],"y=",y[i])
  }
}
