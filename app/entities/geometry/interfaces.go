package geometry




type Point interface {
	Distance(p PointImpl) float64
	IsInline(points ...PointImpl) bool
}
