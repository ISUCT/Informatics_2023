package lab5

import "fmt"

type mouse struct {
 	age    int
 	color  string
 	name   string
 	gender  string
}

func (c mouse) Getage() int {
 return c.age
}

func (c *mouse) Setage(age int) int {
 c.age = age
 return c.age
}
func (c mouse) Getcolor() string {
 return c.color
}
func (c *mouse) Setage(age int) error {
 if age <= 0 || age >= 3 {
  fmt.Printf("Возраст должен быть больше 0 и меньше 3\n")
  return nil
 }
 c.age = age
 return nil
}
func (c mouse) Getname() string {
 return c.name
}
func (c *mouse) Setname(name string) string {
 c.name = name
 return c.name
}
func (c mouse) Getgender() string {
 return c.gender
}

func (c *mouse) Setgender(gender string) string {
 c.gender = gender
 return c.gender
}
func Mymouse(age int, color, name, gender string) mouse {
 return mouse{age, color, name, gender}
}