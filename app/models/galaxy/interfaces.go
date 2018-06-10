package galaxy

import "github.com/juaguz/ml-test/app/models/geometry"

type Planet interface {
	AngularPosition(day uint) uint
	AngularToPoint(day uint) geometry.Point
}
