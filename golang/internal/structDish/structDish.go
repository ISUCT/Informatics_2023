package sctructDish

type bludo struct {
	price   uint8
	calorie float64
	name    string
}

func (b bludo) GetPrice() uint8 {
	return b.price
}

func (b *bludo) ChangeName(name string) string {
	b.name = name
	return b.name
}

func (b *bludo) ChangeCalory(calory float64) float64 {
	b.calorie = calory
	return b.calorie
}

func (b *bludo) ChangePrice(price uint8) uint8 {
	b.price = price
	return b.price
}

func NewDish(price uint8, calory float64, name string) bludo {
	return bludo{
		price:   price,
		calorie: calory,
		name:    name,
	}
}
