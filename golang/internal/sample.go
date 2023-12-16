package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

type Character struct {
	Name  string
	Class string
	Level int
	HP    int
}

func (character *Character) LevelUp(Level int) int {
	character.Level = Level + 1
	return character.Level
}

func (character *Character) ClassChange(RequiredClass string) string {
	character.Class = RequiredClass
	return character.Class
}

func (character *Character) HpChange(CharacterLVL int) int {
	character.HP = character.Level * 10
	return character.HP
}

func Cycle(xn, xk, xd float64) []float64 {
	var numbersCycle = make([]float64, 0, int(((xk-xn)/xd)+1))
	for i := 1.81; i <= 5.31; i = i + float64(0.7) {
		numbersCycle = append(numbersCycle, func_resolution(i))
	}
	return numbersCycle
}

func Array(numbers []float64) []float64 {
	var numbersArray = make([]float64, 0, len(numbers))
	for _, val := range numbers {
		numbersArray = append(numbersArray, func_resolution(val))
	}
	return numbersArray
}

func func_resolution(x float64) float64 {
	var a float64 = 7.2
	var b float64 = 4.2
	r := b * x
	diff := a - r
	mod := math.Abs(diff)
	pow := math.Pow(math.Log10(x), 3)
	fraction := mod / pow
	square := math.Cbrt(fraction)
	y := square

	return y
}
