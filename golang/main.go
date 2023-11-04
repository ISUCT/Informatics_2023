package main

import(
    "fmt"
    "math"
)

func taskA(xn float64, xk float64) {
    var a1 float64 = 4.1
    var b1 float64 = 2.7
    var deltax float64 = 0.4
    var answers1 = []float64{}
    for x:= xn; x<=xk; x+=deltax {
        var y float64 = (a1 * math.Pow(x, 1/3) - b1 * (math.Log(x)/math.Log(5))) / math.Pow(math.Log10(x-1),3)
        answers1 = append(answers1, y)
    }
    fmt.Println(answers1)
}

func taskB() {
    var a2 float64 = 4.1
    var b2 float64 = 2.7
    var znachX = [5]float64{1.9, 2.15, 2.34, 2.74, 3.16}
    var answers2 = []float64{}
    for i:=0; i < len(znachX); i++ {
        var x float64 = znachX[i]
        var y float64 = (a2 * math.Pow(x, 1/3) - b2 * (math.Log(x)/math.Log(5))) / math.Pow(math.Log10(x-1),3)
        answers2 = append(answers2, y)
    }
    fmt.Println(answers2)
}
func main() {
    taskA(1.5,3.5)
    taskB()
}
