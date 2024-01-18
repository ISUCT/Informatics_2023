package struct2

import (
	"fmt"
)

type Mouse struct {
	Brand    string
	Buttons  int
	Wireless bool
}

func (m *Mouse) GetBrand() (string, error) {
	if m.Brand == "" {
		return "", fmt.Errorf("бренд мыши не установлен")
	}
	return m.Brand, nil
}

func (m *Mouse) SetButtons(num int) error {
	if num < 0 {
		return fmt.Errorf("количество кнопок не может быть отрицательным")
	}
	m.Buttons = num
	return nil
}

func (m *Mouse) IsWireless() bool {
	return m.Wireless
}
