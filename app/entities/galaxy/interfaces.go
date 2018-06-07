package galaxy

import "github.com/juaguz/ml-test/app/entities/geometry"

type Planet interface {
	AngularPosition(day uint) uint
	AngularToPoint(day uint) geometry.Point
}
