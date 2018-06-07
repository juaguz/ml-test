package main

import (
	"github.com/juaguz/ml-test/app/entities/galaxy"
	"github.com/juaguz/ml-test/app/entities/geometry"
	"fmt"
)

func main(){
	WeatherPrediction()
}

func WeatherPrediction() {
	ferengi := galaxy.PlanetImpl{Radius: 500, AngularSpeed: 1, ClockWise: true}
	vulcano := galaxy.PlanetImpl{Radius: 1000, AngularSpeed: 5}
	betasoide := galaxy.PlanetImpl{Radius: 2000, AngularSpeed: 3, ClockWise: true}
	sun := geometry.PointImpl{}

	var rain, drought, optimalCondition, maxRainyDay, nothing  uint
	var maxPerimeter float64
	var day uint

	for day = 0; day < 10*360 ; day++ {
		p1 := ferengi.AngularToPoint(day)
		p2 := vulcano.AngularToPoint(day)
		p3 := betasoide.AngularToPoint(day)

		if p1.IsInline(p2, p3) {
			if sun.IsInline(p1, p2, p3) {
				drought += 1
			} else {
				optimalCondition += 1

			}
			continue
		}

		triangle := geometry.Triangle{V1: p1, V2: p2, V3: p3}
		if triangle.PointInside(sun) {
			trianglePerimeter := triangle.Perimeter()
			if maxPerimeter < trianglePerimeter {
				maxRainyDay = day
				maxPerimeter = trianglePerimeter
			}
			rain += 1
			continue
		}
		nothing += 1

	}

	fmt.Printf("Habra %v peridos de sequia \n", drought)
	fmt.Printf("Habra %v peridos de lluvia y el pico en el dia %v \n", rain, maxRainyDay)
	fmt.Printf("Habra %v peridos de condiciones optimas de clima y temperatura \n", optimalCondition)
	fmt.Printf("Habra %v peridos de nada \n", nothing)

}
