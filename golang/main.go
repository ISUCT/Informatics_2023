package main

import (
	"fmt"
	"math"
)

func main(){
	a:= 2.0
	b:= 3.0
	n:= 0.11
	k:= 0.36
	d:= 0.05 


for x := n; x <= k; x += d {
	y:= math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))
	fmt.Printf("x = %.2f, y = %.2f\n", x, y)
}
}
