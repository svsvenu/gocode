package inheritance

type ShapeOperatons interface {
	GetArea() int
	GetLength() int
}

type Circle struct {
	radius int
}

func (c *Circle) GetArea() int {
	return c.radius * c.radius
}
func (c *Circle) GetLength() int {
	return 2 * c.radius
}

type Rectange struct {
	length int
	width  int
}

func (r Rectange) GetArea() int {
	return r.length * r.width
}
func (r Rectange) GetLength() int {
	return 2*r.length + 2*r.width
}
