package main

import (
	"errors"
	"fmt"
)

type GameCharacter struct {
	class    string
	level    int
	nickname string
}

func NewGameCharacter(class string, level int, nickname string) (*GameCharacter, error) {
	if level < 0 {
		return nil, errors.New("выйди и удали игру")
	}

	return &GameCharacter{class, level, nickname}, nil
}

func main() {
	character, err := NewGameCharacter("Warrior", 10, "Player1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Class: %s, Level: %d, Nickname: %s\n", character.class, character.level, character.nickname)
}
