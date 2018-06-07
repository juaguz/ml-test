package entities


type PointImpl struct {
	X,Y float64
}

func (pi PointImpl) Distance(p Point) uint {
	return uint(1)
}

