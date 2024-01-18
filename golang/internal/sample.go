package internal

import (
	"fmt"
)

type Auto struct {
	speed    int64
	color    string
	size     string
	distance int64
}

func (m *Auto) Getspeed() int64 {
	if m.speed <= 0 {
		fmt.Println("Ты физику уважаешь? Отрицательная скорость? Цвет и размер уже не важны.")
		m.speed = 0
	} else if m.speed > 300 {
		fmt.Println("Это птица? Это самолёт? Нет, это машину в воздух хотят запустить. Цвет и размер уже не уважны.")
		m.speed = 0
	}
	return m.speed
}
func (m *Auto) Setspeed(speed int64) {
	m.speed = speed
}
func (m *Auto) Getcolor() string {
	return m.color
}
func (m *Auto) Setcolor(color string) {
	m.color = color
}
func (m *Auto) Setdistance(distance int64) {
	m.distance = distance
}
func (m *Auto) Getdistance() int64 {
	if m.distance < 0 {
		fmt.Println("Ну я так не играю, давай нормальное расстояние")
		m.distance = 0
	}
	return m.distance
}
func (m *Auto) Getsize() string {
	return m.size
}
func (m *Auto) Setsize(size string) {
	m.size = size
}
func calculateTime(m *Auto) float64 {
	distance := float64(m.Getdistance())
	speed := float64(m.Getspeed())

	if speed == 0 {
		return 0
	}

	time := distance / speed
	if time == 1 {
		time = 0
	}
	return time
}
func main() {
	automobile := Auto{}
	fmt.Println("Выберите цвет машины:")
	fmt.Println("1. Красный")
	fmt.Println("2. Синий")
	fmt.Println("3. Зеленый")
	fmt.Println("4. Розовый")

	var ColorChoice int
	fmt.Scanln(&ColorChoice)

	for {
		switch ColorChoice {
		case 1:
			automobile.Setcolor("Красный")
			break
		case 2:
			automobile.Setcolor("Синий")
			break
		case 3:
			automobile.Setcolor("Зеленый")
			break
		case 4:
			automobile.Setcolor("Розовый")
			break
		default:
			fmt.Println("Ошибка ввода, неверный выбор")
			fmt.Println("Выберите цвет машины:")
			fmt.Println("1. Красный")
			fmt.Println("2. Синий")
			fmt.Println("3. Зеленый")
			fmt.Println("4. Розовый")
			fmt.Scanln(&ColorChoice)
			continue
		}
		break
	}
	fmt.Println("Выберите размер машины:")
	fmt.Println("1. Двуместная")
	fmt.Println("2. Четырехместная")
	var SizeChoice int
	fmt.Scanln(&SizeChoice)
	for {
		switch SizeChoice {
		case 1:
			automobile.Setsize("Двуместная")
			break
		case 2:
			automobile.Setsize("Четырехместная")
			break
		default:
			fmt.Println("Ошибка ввода, неверный выбор")
			fmt.Println("Выберите цвет машины:")
			fmt.Println("1. Двуместная")
			fmt.Println("2. Четырехместная")
			fmt.Scanln(&SizeChoice)
			continue
		}
		break
	}
	var selectspeed int64
	fmt.Println("С какой скоростью едет машина:")
	fmt.Scanln(&selectspeed)
	automobile.Setspeed(selectspeed)
	err := automobile.Getspeed()
	if err != 0 {
		fmt.Println("Цвет автомобиля:", automobile.Getcolor())
		fmt.Println("Размер автомобиля:", automobile.Getsize())
		fmt.Println("текущая скорость:", automobile.Getspeed(), "км/ч")
	} else {
		fmt.Println("")
	}
	fmt.Println("Хотите посмотреть за какое время ваша машина проедет любое расстояние?")
	fmt.Println("1. Да, конечно!!!")
	fmt.Println("2. Нет, спасибо")

	var DistanceChoice int
	fmt.Scanln(&DistanceChoice)

	var Dist string
	switch DistanceChoice {
	case 1:
		Dist = "Да, конечно!!!"
	case 2:
		Dist = "Нет, спасибо"
	default:
		fmt.Println("Ошибка ввода, неверный выбор")

	}
	if Dist == "Да, конечно!!!" {
		fmt.Println("Отлично, введите расстояние в километрах:")
		var Dista int64
		fmt.Scanln(&Dista)
		automobile.Setdistance(Dista)
		err := automobile.Getdistance()
		for {
			if err != 0 {
				time := calculateTime(&automobile)
				fmt.Println("Время требуемое для преодоления расстояния", time, "часов")
				break
			} else {
				fmt.Scanln(&Dista)
				time := calculateTime(&automobile)
				fmt.Println("Время требуемое для преодоления расстояния", time, "часов")
				break
			}
		}
	}
}
