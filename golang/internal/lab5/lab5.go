package lab5

import (
	"errors"
)

type ACE struct {
	firepower float64
	speed     float64
	armor     float64
	name      string
}

func NewACE(firepower float64, speed float64, armor float64, name string) (*ACE, error) {
	p := &ACE{
		firepower: firepower,
		speed:     speed,
		armor:     armor,
		name:      name,
	}
	if err := p.SetFirepower(firepower); err != nil {
		return nil, err
	}
	if err := p.SetSpeed(speed); err != nil {
		return nil, err
	}
	if err := p.SetArmor(armor); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *ACE) GetFirepower() float64 {
	return p.firepower
}

func (p *ACE) SetFirepower(firepower float64) error {
	if firepower <= 0 || firepower > 5000 {
		return errors.New("неверно указана огневая мощь, пожалуйста, измените параметр")
	}
	p.firepower = firepower
	return nil
}

func (p *ACE) GetSpeed() float64 {
	return p.speed
}

func (p *ACE) SetSpeed(speed float64) error {
	if speed <= 0 || speed > 5000 {
		return errors.New("неверно указана скорость, пожалуйста, измените параметр")
	}
	p.speed = speed
	return nil
}

func (p *ACE) GetArmor() float64 {
	return p.armor
}

func (p *ACE) SetArmor(armor float64) error {
	if armor <= 0 || armor > 5000 {
		return errors.New("неверно указано бронирование, пожалуйста, измените параметр")
	}
	p.armor = armor
	return nil
}

func (p *ACE) GetName() string {
	return p.name
}

func (p *ACE) SetName(name string) {
	p.name = name
}