package lab5

type table struct {
	length float64
	width  float64
	height float64
}

func (t table) GetLenght() float64 {
	return t.length
}

func (t table) GetWidht() float64 {
	return t.width
}

func (t table) GetHeight() float64 {
	return t.height
}

func The_Table(table_lenght, table_widht, table_height float64) table {
	return table{
		length: table_lenght,
		width:  table_widht,
		height: table_height,
	}
}
