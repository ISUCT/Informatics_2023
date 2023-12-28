package main

import (
    "fmt"
  
    "isuct.ru/informatics2022/internal/lab5"
    //  "isuct.ru/informatics2022/internal/lab4"
)

func main(){
    // lab4
    // fmt.Println("Фёдорова Ольга Андреевна")
    // fmt.Println("Задача A")

    // x,y:=lab4.TaskA(1.2, 4.2, 0,6)
    // for i := range y{
    //   fmt.Println("При x=", x[i],"y=",y[i])
    // }
    // fmt.Println("Задача B")
    // var list []float64=[]float64{1.16, 1.32, 1.47, 1.65, 1.93}
    // x,y=lab4.TaskB(list)
    // for i := range y{
    //   fmt.Println("При x=",x[i],"y=",y[i])
    // }
    mouse1 := lab5.Mymouse(2, "чёрный", "Вероника", "женский")
    mouse2 := lab5.Mymouse(1, "белый", "Оля", "женский")
    fmt.Printf("возраст: %d\n", mouse1.Getage())
    fmt.Printf("цвет: %s\n", mouse1.Getcolor())
    fmt.Printf("имя %s\n", mouse1.Getname())
    fmt.Printf("пол %s\n", mouse1.Getgender())
    mouse2.Setage(2)
    mouse2.Setage(7)
    mouse2.Setgeder("женский")
    fmt.Printf("возраст: %s\n", mouse2.Getage())
    fmt.Printf("пол %s\n", mouse2.Getmileage())
}

