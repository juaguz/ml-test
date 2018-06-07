package entities

import (
	"math"
	"github.com/juaguz/ml-test/app/utils"
)


type PlanetImpl struct {
	Radius uint
	AngularSpeed uint
	ClockWise bool
}

func (p *PlanetImpl) AngularPosition(day uint) float64 {
	var angularPosition float64

	angularPosition = float64(day * p.AngularSpeed)

	if p.ClockWise {
		return angularPosition
	}

	return 360-angularPosition

}
func (p *PlanetImpl) AngularToPoint(day uint) Point {
	degree := p.AngularPosition(day)
	radians := utils.DegreeToRadian(degree)
	x := math.Round(math.Cos(radians) * float64(p.Radius))
	y := math.Round(math.Sin(radians) * float64(p.Radius))
	return PointImpl{x, y}
}
