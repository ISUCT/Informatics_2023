package main

import (
	"bufio"
	"fmt"
	"os"

	"isuct.ru/informatics2022/internal"
)

func DefaultConsts() internal.FormulaConsts {
	return internal.FormulaConsts{A: 2.0, B: 3.0}
}

func DefaultDeltaInfo() internal.DeltaInfo {
	return internal.DeltaInfo{XStart: 0.11, XEnd: 0.36, XStep: 0.05}
}

var Xes = []float64{0.08, 0.26, 0.37, 0.48, 0.56}

func main() {
	stepResult := internal.CalculateResultsWithStep(DefaultConsts(), DefaultDeltaInfo())
	xesResult := internal.CalculateResultWithSlices(DefaultConsts(), Xes)

	fmt.Println("With steps:")
	internal.PrintSlice(stepResult)

	fmt.Println("With const set:")
	internal.PrintSlice(xesResult)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
