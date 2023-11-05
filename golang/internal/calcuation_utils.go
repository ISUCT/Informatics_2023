package internal

import "math"

type FormulaConsts struct {
	A float64
	B float64
}

type DeltaInfo struct {
	XStart float64
	XEnd   float64
	XStep  float64
}

type ResultRecord struct {
	X float64
	Y float64
}

func (info DeltaInfo) getPossibleStepCount() int {
	return int((info.XEnd-info.XStart)/info.XStep) + 1
}

func Formula(x float64, a float64, b float64) ResultRecord {
	return ResultRecord{X: x, Y: math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))}
}

func CalculateResultsWithStep(consts FormulaConsts, deltaInfo DeltaInfo) []ResultRecord {
	results := make([]ResultRecord, 0, deltaInfo.getPossibleStepCount())
	for x := deltaInfo.XStart; x <= deltaInfo.XEnd; x += deltaInfo.XStep {
		results = append(results, Formula(x, consts.A, consts.B))
	}

	return results
}

func CalculateResultWithSlices(consts FormulaConsts, xes []float64) []ResultRecord {
	results := make([]ResultRecord, 0, len(xes))
	for _, x := range xes {
		results = append(results, Formula(x, consts.A, consts.B))
	}
	return results
}
