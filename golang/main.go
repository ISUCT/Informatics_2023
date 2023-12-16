package main 
 
import ( 
 "fmt" 
 "isuct.ru/informatics2022/internal" 
) 
 
 
func main() { 
 fmt.Println("Кувшинов Владислав") 
 fmt.Println(internal.CalculateFunctionInRanges(0.26, 0.66, 0.08)) 
 fmt.Println(internal.CalculateFunctionWithXValues([]float64{0.1, 0.35, 0.4, 0.55, 0.6})) 
}

