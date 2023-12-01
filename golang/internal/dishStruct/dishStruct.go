package dishstruct

import (
	"fmt"
)

type dish struct {
	price int
	vid   string
	name  string
}

func NewDish(setPrice int, setVid, setName string) dish {
	return dish{
		price: setPrice,
		vid:   setVid,
		name:  setName,
	}
}

func (d dish) GetPrice() int {
	return d.price
}

func (d dish) GetVid() string {
	return d.vid
}

func (d dish) GetName() string {
	return d.name
}

func (d *dish) SetName(name string) {
	d.name = name
}

func (d *dish) SetPrice(price int) error {
	if price <= 5000 {
		d.price = price
		return nil
	}
	return fmt.Errorf("The dish is not affordable")
}

func (d *dish) TipTheWaiter() {
	fmt.Println("Thx for your tip!")
}
