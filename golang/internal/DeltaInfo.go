package internal

type DeltaInfo struct {
	xStart float64
	xEnd   float64
	xStep  float64
}

func (self *DeltaInfo) XStart() float64 {
	return self.xStart
}

func (self *DeltaInfo) SetXStart(value float64) *DeltaInfo {
	self.xStart = value
	return self
}

func (self *DeltaInfo) XEnd() float64 {
	return self.xEnd
}

func (self *DeltaInfo) SetXEnd(value float64) *DeltaInfo {
	self.xEnd = value
	return self
}

func (self *DeltaInfo) XStep() float64 {
	return self.xStep
}

func (self *DeltaInfo) SetXStep(value float64) *DeltaInfo {
	self.xStep = value
	return self
}

func NewDeltaInfo() DeltaInfo {
	return DeltaInfo{xStart: 0.11, xEnd: 0.36, xStep: 0.05}
}
