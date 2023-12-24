package structures

type character struct {
	class string
	lvl   uint32
	name  string
}

func (c character) Name() string {
	return c.name
}
func (c *character) Class(class string) string {
	c.class = class
	return c.class
}
func (c *character) Lvl(lvl uint32) uint32 {
	c.lvl = lvl
	return c.lvl
}
func Char(name, class string, lvl uint32) character {
	return character{
		name:  name,
		class: class,
		lvl:   lvl,
	}
}
