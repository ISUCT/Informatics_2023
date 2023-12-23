package lab5

import (
	"fmt"
)

type table struct {
	length float64
	width  float64
	height float64
}

func (t table) GetLength() float64 {
	return t.length
}

func (t table) GetWidth() float64 {
	return t.width
}

func (t table) GetHeight() float64 {
	return t.height
}

func NewTable(checkLength, checkWidth, checkHeight float64) (table, error) {
	var t table
	err := t.SetLength(checkLength)
	if err != nil {
		return t, err
	}
	err = t.SetWidth(checkWidth)
	if err != nil {
		return t, err
	}
	err = t.SetHeight(checkHeight)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (t *table) SetLength(length float64) error {
	if length > 0 {
		t.length = length
		return nil
	}
	return fmt.Errorf("invalid table length")
}

func (t *table) SetWidth(width float64) error {
	if width > 0 {
		t.width = width
		return nil
	}
	return fmt.Errorf("invalid table width")
}

func (t *table) SetHeight(height float64) error {
	if height > 0 {
		t.height = height
		return nil
	}
	return fmt.Errorf("invalid table height")
}
