package internal

type FormulaConsts struct {
	a float64
	b float64
}

func (self *FormulaConsts) A() float64 {
	return self.a
}

func (self *FormulaConsts) SetA(value float64) *FormulaConsts {
	self.a = value
	return self
}

func (self *FormulaConsts) B() float64 {
	return self.b
}

func (self *FormulaConsts) SetB(value float64) *FormulaConsts {
	self.b = value
	return self
}

func NewConsts() FormulaConsts {
	return FormulaConsts{a: 2.0, b: 3.0}
}
