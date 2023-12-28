package lab5

import "fmt"

type car struct {
	model    string
	maxspeed int
	price    int
	mileage  int
}

func (c car) Getmodel() string {
	return c.model
}

func (c *car) Setmodel(model string) string {
	c.model = model
	return c.model
}
func (c car) Getmaxspeed() int {
	return c.maxspeed
}
func (c *car) Setmaxspeed(maxspeed int) error {
	if maxspeed <= 0 || maxspeed >= 420 {
		fmt.Printf("Максимальная скорость должна быть больше 0 и меньше 420\n")
		return nil
	}
	c.maxspeed = maxspeed
	return nil
}
func (c car) Getprice() int {
	return c.price
}
func (c *car) Setprice(price int) int {
	c.price = price
	return c.price
}
func (c car) Getmileage() int {
	return c.mileage
}

func (c *car) Setmileage(mileage int) int {
	c.mileage = mileage
	return c.mileage
}
func Mycar(model string, maxspeed, price, mileage int) car {
	return car{model, maxspeed, price, mileage}
}
