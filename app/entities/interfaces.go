package entities


type Planet interface {
	AngularPosition(day uint) uint
	AngularToPoint(day uint) Point
}

type Point interface {
	Distance(p PointImpl) float64
}
