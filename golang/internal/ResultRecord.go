package internal

type ResultRecord struct {
	x float64
	y float64
}

func (self *ResultRecord) X() float64 {
	return self.x
}

func (self *ResultRecord) Y() float64 {
	return self.y
}
