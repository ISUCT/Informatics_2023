package lab4

import "errors"

const fieldWidht = 8
const fieldHeight = 8
const maxShipLength = 4

type ship struct {
	width  uint8
	height uint8
	x      uint8
	y      uint8
}

func NewShip(width, height, x, y uint8) (*ship, error) {

	if width > 1 && height > 1 {
		return nil, errors.New("Unsupported ship sizes.")
	}

	if width == 0 || height == 0 {
		return nil, errors.New("Unsupported ship sizes.")
	}

	if width > maxShipLength || height > maxShipLength {
		return nil, errors.New("Unsupported ship sizes.")
	}

	ship := ship{
		width:  width,
		height: height,
	}

	error := ship.SetX(x)
	if error != nil {
		return nil, error
	}

	error = ship.SetY(y)
	if error != nil {
		return nil, error
	}

	return &ship, nil
}

func (instance *ship) Width() uint8 {
	return instance.width
}

func (instance *ship) Height() uint8 {
	return instance.height
}

func (instance *ship) X() uint8 {
	return instance.x
}

func (instance *ship) Y() uint8 {
	return instance.y
}

func (instance *ship) SetX(x uint8) error {
	if x+instance.width > fieldWidht {
		return errors.New("Ship go out of field bounds.")
	}

	instance.x = x
	return nil
}

func (instance *ship) SetY(y uint8) error {
	if y+instance.height > fieldWidht {
		return errors.New("Ship go out of field bounds.")
	}

	instance.y = y
	return nil
}
