package rect

type Rect struct {
	X int
	Y int
}

func (r Rect) GetArea() int {
	return r.X * r.Y
}

func (r Rect) GetCircuit() int {
	return 2*r.X + 2*r.Y
}
