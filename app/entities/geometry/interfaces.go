package geometry




type Point interface {
	Distance(p PointImpl) float64
}
