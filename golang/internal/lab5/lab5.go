package lab5

import (
	"errors"
	"fmt"
)

type table struct {
	height int
	length int
	width  int
}

func NewTable(height, length, width int) (*table, error) {
	t := &table{
		height: height,
		length: length,
		width:  width,
	}

	if err := t.SetHeight(height); err != nil {
		return nil, err
	}
	if err := t.SetLength(length); err != nil {
		return nil, err
	}
	if err := t.SetWidth(width); err != nil {
		return nil, err
	}

	return t, nil
}

func (t table) GetHeight() int {
	return t.height
}
func (t table) GetLength() int {
	return t.length
}
func (t table) GetWidth() int {
	return t.width
}
func (t *table) SetHeight(height int) error {
	if height >= 0 && height <= 120 {
		t.height = height
		return nil
	}
	return errors.New("ErrIncorrectHeight")
}

func (t *table) SetLength(length int) error {
	if length >= 0 && length <= 140 {
		t.length = length
		return nil
	}
	return errors.New("ErrIncorrectLength")
}

func (t *table) SetWidth(width int) error {
	if width >= 0 && width <= 60 {
		t.width = width
		return nil
	}
	return errors.New("ErrIncorrectWidth")
}

func (t *table) GetTablesSize() {
	fmt.Println("Here are the dimensions of the table!")
}
