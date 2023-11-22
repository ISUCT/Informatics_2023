package main

import (
	"bufio"
	"fmt"
	"os"

	"isuct.ru/informatics2022/internal"
)

func main() {
	stepResult := internal.CalculateResultsWithStep(internal.NewConsts(), internal.NewDeltaInfo())
	xesResult := internal.CalculateResultWithSlices(internal.NewConsts(), internal.DefaultXes())

	fmt.Println("With steps:")
	internal.PrintSlice(stepResult)

	fmt.Println("With const set:")
	internal.PrintSlice(xesResult)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
