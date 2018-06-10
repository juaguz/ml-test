package galaxy

import (
	"testing"
	"github.com/juaguz/ml-test/app/models/geometry"
)

func TestAngularPosition(t *testing.T) {
	expected := float64(1)
	p := PlanetImpl{
		Radius: uint(500),
		AngularSpeed: uint(1),
		ClockWise:true,
	}
	angularPosition := p.AngularPosition(1)
	if expected != angularPosition {
		t.Fatalf("the angular position should be %v not %v",expected,angularPosition)
	}
}


func TestAngularPositionAntiClockWise(t *testing.T) {
	expected := float64(355)
	p := PlanetImpl{
		Radius: uint(1000),
		AngularSpeed: uint(5),
		ClockWise:false,
	}
	angularPosition := p.AngularPosition(1)
	if expected != angularPosition {
		t.Fatalf("the angular position should be %v not %v",expected,angularPosition)
	}
}

func TestPlanetImplAngularToPoint(t *testing.T) {
	expected := geometry.PointImpl{996, -87}
	p := PlanetImpl{
		Radius: uint(1000),
		AngularSpeed: uint(5),
		ClockWise:false,
	}
	point := p.AngularToPoint(1)
	if point != expected {
		t.Fatalf("the angular position should be %v not %v",expected, point)
	}
}