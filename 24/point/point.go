package point

type Point struct {
	X, Y float64
}

type Line struct {
	Point
}

func New(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance() float64 {
	if p.X > p.Y {
		return p.X - p.Y
	} else {
		return p.Y - p.X
	}
}
