package lab5

import "fmt"

type ship struct {
	length       float64
	displacement float64
	maxspeed     int8
}

func (s ship) GetLength() float64 {
	return s.length
}

func (s ship) GetDisplacement() float64 {
	return s.displacement
}

func (s ship) GetMaxSpeed() int8 {
	return s.maxspeed
}
func TheShip(lengthOfShip, displacementOfShip, maxspeedOfShip float64) (ship, error) {
	var s ship
	err := s.SetLength(lengthOfShip)
	if err != nil {
		return s, err
	}
	err = s.SetDisplacement(displacementOfShip)
	if err != nil {
		return s, err
	}
	err = s.SetMasxSpeed(int8(maxspeedOfShip))
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s *ship) SetLength(length float64) error {
	if length > 0 {
		s.length = length
		return nil
	}
	return fmt.Errorf("the wrong length of the ship")
}

func (s *ship) SetDisplacement(displacement float64) error {
	if displacement > 0 {
		s.displacement = displacement
		return nil
	}
	return fmt.Errorf("the wrong displacement of the ship")
}

func (s *ship) SetMasxSpeed(maxspeed int8) error {
	if maxspeed > 0 {
		s.maxspeed = maxspeed
		return nil
	}
	return fmt.Errorf("the wrong max speed of the ship")
}