package dishstruct

import (
	"errors"
	"fmt"
)

type dish struct {
	price int
	vid   string
	name  string
}

func NewDish(price int, setVid, setName string) (*dish, error) {
	d := &dish{
		price: price,
		vid:   setVid,
		name:  setName,
	}
	if err := d.SetPrice(price); err != nil {
		return nil, err
	}
	return d, nil
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
	if price >= 5000 || price < 0 {
		return errors.New("The dish is not affordable")
	}
	d.price = price
	return nil
}

func (d *dish) TipTheWaiter() {
	fmt.Println("If you liked the service you can leave a tip!")
}
