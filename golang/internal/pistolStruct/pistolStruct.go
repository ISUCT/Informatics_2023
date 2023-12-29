package pistolStruct

import "fmt"

type pistol struct {
	magazineCapacity int
	model            string
	caliber          string
}

func (c pistol) GetmagazineCapacity() int {
	return c.magazineCapacity
}

func NewPistol(capacityVarious int, modelVarious, caliberVarious string) (pistol, error) {
	var p pistol = pistol{
		model:   modelVarious,
		caliber: caliberVarious,
	}
	var err = p.SetmagazineCapacity(capacityVarious)
	return p, err
}

func (c *pistol) SetmagazineCapacity(capacity int) error {
	if capacity >= 5 && capacity <= 30 {
		c.magazineCapacity = capacity
		return nil
	}
	return fmt.Errorf("invalid pistol magazine capacity")
}

func (c pistol) GetModel() string {
	return c.model
}

func (c *pistol) SetModel(model string) {
	c.model = model
}

func (c pistol) GetCaliber() string {
	return c.caliber
}
