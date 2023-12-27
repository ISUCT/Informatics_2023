package lab5

import (
	"fmt"
)

type tv struct {
	model string
	color string
	sound int
}

func NewTv(tv_model, tv_color string, tv_sound int) (tv, error) {
	var t tv = tv{
		model: tv_model,
	}
	t.SetColor(tv_color)
	var err = t.SetSound(tv_sound)
	return t, err
}

func (t *tv) TurnOn() {
	fmt.Printf("%s включился", t.GetModel())
}

func (t *tv) SetModel(model string) {
	t.model = model
}

func (t *tv) SetColor(color string) {
	t.color = color
}

func (t *tv) SetSound(sound int) error {
	if sound <= 100 && sound >= 0 {
		t.sound = sound
		return nil
	} else {
		return fmt.Errorf("Громкость может быть от 0 до 100")
	}
}

func (t *tv) GetModel() string {
	return t.model
}

func (t *tv) GetColor() string {
	return t.color
}

func (t *tv) GetSound() int {
	return t.sound
}
