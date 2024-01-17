package character

import "fmt"

type Character struct {
	Name      string
	Intellect int
	Psyche    int
	Physique  int
	Motorics  int
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) GetOverallPoints() int {
	return c.Intellect + c.Psyche + c.Physique + c.Motorics
}

func (c *Character) PrintDetails() {
	fmt.Printf("Name: %s\nIntellect: %d\nPsyche: %d\nPhysique: %d\nMotorics: %d\n", c.Name, c.Intellect, c.Psyche, c.Physique, c.Motorics)
}
