package main

import "fmt"

type Television struct {
	Brand   string
	Model   string
	Channel int
}

func (tv *Television) SetChannel(channel int) {
	tv.Channel = channel
}

func (tv *Television) GetChannel() int {
	return tv.Channel
}

func (tv *Television) Display() {
	fmt.Printf("Телевизор %s %s находится на канале %d\n", tv.Brand, tv.Model, tv.Channel)
}

func main() {
	tv := Television{
		Brand:   "Samsung",
		Model:   "QLED",
		Channel: 1,
	}

	tv.Display()
	tv.SetChannel(5)
	tv.Display()
}
