package plane

import "fmt"

type plane struct {
	speed int
	model string
	color string
}

func NewPlane(model_name, color string, speed int) (plane, error) {
	var p plane = plane{
		model: model_name,
	}
	var err = p.SetSpeed(speed)
	p.SetColor(color)
	return p, err
}

func (p *plane) SetSpeed(plane_speed int) error {
	if plane_speed <= 3529 && plane_speed >= 0 {
		p.speed = plane_speed
		return nil
	} else {
		return fmt.Errorf("Самолет не может лететь c такой скоростью")
	}
}

func (p *plane) GetSpeed() int {
	return p.speed
}

func (p *plane) SetColor(plane_color string) {
	p.color = plane_color
}

func (p *plane) GetColor() string {
	return p.color
}

func (p *plane) GetModel() string {
	return p.model
}
