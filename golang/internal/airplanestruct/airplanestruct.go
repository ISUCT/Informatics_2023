package airplanestruct

import "fmt"

type airplane struct {
	speed        int
	model        string
	manufcountry string
}

func (a airplane) GetSpeed() int {
	return a.speed
}

func NewAirplane(speedVar int, modelVar, manufcountryVar string) (airplane, error) {
	var a airplane = airplane{
		model:        modelVar,
		manufcountry: manufcountryVar,
	}
	var err = a.SetSpeed(speedVar)
	return a, err
}

func (a *airplane) SetSpeed(speed int) error {
	if speed > 0 && speed < 10000 {
		a.speed = speed
		return nil
	}
	return fmt.Errorf("failed SetSpeed() for airplane model \"%s\" - invalid speed", a.GetModel())
}

func (a airplane) GetModel() string {
	return a.model
}

func (a *airplane) SetModel(model string) {
	a.model = model
}

func (a airplane) GetManufcountry() string {
	return a.manufcountry
}

func (a *airplane) SetManufcountry(manufcountry string) {
	a.manufcountry = manufcountry
}
