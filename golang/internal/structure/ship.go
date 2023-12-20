package ship

import "fmt"

type ship struct {
	length        float64
	displacement  float64
	maxSpeedKnots float64
}

func NewShip(length float64, displacement float64, maxSpeedKnots float64) (ship, error) {
	var s ship = ship{
		maxSpeedKnots: maxSpeedKnots,
	}
	s.SetLenght(length)
	var err = s.SetDisplacement(displacement)
	return s, err
}

func (s *ship) SetMaxSpeedKnots(maxSpeedKnots float64) {
	s.maxSpeedKnots = maxSpeedKnots
}

func (s *ship) SetDisplacement(displacement float64) error {
	if displacement > 0 {
		s.displacement = displacement
		return nil
	} else {
		return fmt.Errorf("Объём не может быть отрицательным!")
	}
}

func (s *ship) SetLenght(length float64) error {
	if length > 0 {
		s.length = length
		return nil
	} else {
		return fmt.Errorf("Длина не может быть отрицательной!")
	}
}

func (s *ship) GetLength() float64 {
	return s.length
}

func (s *ship) GetMaxSpeedKnots() float64 {
	return s.maxSpeedKnots
}

func (s *ship) GetDisplacement() float64 {
	return s.displacement
}
