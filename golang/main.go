package main

import (
	"fmt"
	"math"
)

func main() {
	a := 4.1
	b := 2.7
	fmt.Println("13 Вариант")
	fmt.Println("задача А")
	for x := 1.2; x <= 5.2; x += 0.8 {
		res1 := a*math.Sqrt(x) - b*(math.Log(x)/math.Log(5))
		res2 := math.Log10(math.Abs(x - 1))
		res3 := res1 / res2
		fmt.Println(res3)
	}
	fmt.Println("задача B")
	x1 := 1.9
	x2 := 2.15
	x3 := 2.34
	x4 := 2.73
	x5 := 3.16

	y1 := ((a*math.Sqrt(x1) - b*(math.Log(x1)/math.Log(5))) / (math.Log10(math.Abs(x1 - 1))))
	fmt.Println(y1)
	y2 := ((a*math.Sqrt(x2) - b*(math.Log(x2)/math.Log(5))) / (math.Log10(math.Abs(x2 - 1))))
	fmt.Println(y2)
	y3 := ((a*math.Sqrt(x3) - b*(math.Log(x3)/math.Log(5))) / (math.Log10(math.Abs(x3 - 1))))
	fmt.Println(y3)
	y4 := ((a*math.Sqrt(x4) - b*(math.Log(x4)/math.Log(5))) / (math.Log10(math.Abs(x4 - 1))))
	fmt.Println(y4)
	y5 := ((a*math.Sqrt(x5) - b*(math.Log(x5)/math.Log(5))) / (math.Log10(math.Abs(x5 - 1))))
	fmt.Println(y5)
}
