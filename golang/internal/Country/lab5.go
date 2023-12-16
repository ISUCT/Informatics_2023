package country

// Определение структуры страны
type Country struct {
	CountryName string
	Population  int
	Area        float64
	Capital     string
}

// Методы для структуры Country
func (c Country) GetPopulation() int {
	return c.Population
}

func (c Country) GetArea() float64 {
	return c.Area
}
func (c Country) GetCapital() string {
	return c.Capital
}
func (c Country) GetCountryName() string {
	return c.CountryName
}
