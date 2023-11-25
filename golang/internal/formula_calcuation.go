package internal

import (
	"math"
)

func (info DeltaInfo) getPossibleStepCount() int {
	return int((info.xEnd-info.xStart)/info.xStep) + 1
}

func Formula(x float64, a float64, b float64) ResultRecord {
	return ResultRecord{x: x, y: math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))}
}

func CalculateResultsWithStep(consts FormulaConsts, deltaInfo DeltaInfo) []ResultRecord {
	results := make([]ResultRecord, 0, deltaInfo.getPossibleStepCount())
	for x := deltaInfo.xStart; x <= deltaInfo.xEnd; x += deltaInfo.xStep {
		results = append(results, Formula(x, consts.a, consts.b))
	}

	return results
}

func CalculateResultWithSlices(consts FormulaConsts, xes []float64) []ResultRecord {
	results := make([]ResultRecord, 0, len(xes))
	for _, x := range xes {
		results = append(results, Formula(x, consts.a, consts.b))
	}
	return results
}

func DefaultXes() []float64 {
	return []float64{0.08, 0.26, 0.37, 0.48, 0.56}
}
