package internal 
 
import "math" 
 
func Summ(a, b int) int { 
 return a + b 
} 
 
func calculateFunction(x float64) float64 { 
 arcsin := math.Asin(x) 
 arccos := math.Acos(x) 
 return math.Pow( 
  math.Pow(arcsin, 2.0)+ 
   math.Pow(arccos, 4.0), 
  3.0) 
} 
 
func CalculateFunctionInRanges(xStart, xEnd, xStep float64) []float64 { 
 results := make([]float64, 0, int((xEnd - xStart)/xStep + 1.0))
 for x := xStart; x <= xEnd; x += xStep { 
  results = append(results, calculateFunction(x)) 
 } 
 return results 
} 
func CalculateFunctionWithXValues(xes []float64) []float64 { 
 results := make([]float64, 0, len(xes)) 
 for _, x := range xes { 
  results = append(results, calculateFunction(x)) 
 } 
 return results 
}