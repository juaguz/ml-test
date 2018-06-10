package geometry

import (
	"math"
)

type PointImpl struct {
	X,Y float64
}

//Calucla la distancia entre dos pintos
func (pi PointImpl) Distance(p PointImpl) float64 {
	x1 := pi.X
	x2 := p.X
	y1 := pi.Y
	y2 := p.Y
	xPow := math.Pow(x2-x1,2)
	yPow := math.Pow(y2-y1,2)
  	return math.Sqrt(xPow+yPow)
}

//Determina si un conjunto de puntos son colineales
// Por ejemplo los 3 planetas mas el sol
// Los puntos tienen que estar ordenados
//https://vitual.lat/determinar-si-tres-puntos-son-colineales-con-la-formula-de-la-distancia/
func (pi PointImpl) IsCollinear(points ...PointImpl) bool {
	pointsLen := len(points)
	if pointsLen < 1 {
		panic("not enough arguments")
	}
	distance := pi.Distance(points[pointsLen-1])

	var sumDistance float64

	points = append([]PointImpl{pi}, points...)
	for index := range points {
		if index <= pointsLen - 1 {
			sumDistance = sumDistance+points[index].Distance(points[index+1])
		}
	}

	return distance == sumDistance

}