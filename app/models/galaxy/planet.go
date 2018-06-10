package galaxy

import (
	"math"
	"github.com/juaguz/ml-test/app/utils"
	"github.com/juaguz/ml-test/app/models/geometry"
)


type PlanetImpl struct {
	Radius uint
	AngularSpeed uint
	ClockWise bool
}

//Calcula la posicion angular dado un dia
func (p *PlanetImpl) AngularPosition(day uint) float64 {
	var angularPosition float64

	angularPosition = float64(day * p.AngularSpeed)

	if p.ClockWise {
		return angularPosition
	}

	return 360-angularPosition

}
//Transforma la posicion angular a un punto en el eje cartesiano
func (p *PlanetImpl) AngularToPoint(day uint) geometry.PointImpl {
	degree := p.AngularPosition(day)
	radians := utils.DegreeToRadian(degree)
	x := math.Round(math.Cos(radians) * float64(p.Radius))
	y := math.Round(math.Sin(radians) * float64(p.Radius))
	return geometry.PointImpl{X: x, Y: y}
}
