package services

import (
	"github.com/juaguz/ml-test/app/models/galaxy"
	"github.com/juaguz/ml-test/app/models/geometry"
	"fmt"
	"github.com/juaguz/ml-test/app/models/weather"
)

type weatherPrediction struct {
	WeatherService
	SaveDatabase bool
	PrintPrediction bool
}

func (wp *weatherPrediction)save(w weather.Weather) {
	if wp.SaveDatabase {
		wp.WeatherService.Store(&w)
	}
}

func NewWeatherPrediction(service WeatherService, saveDatabase bool,printPrediction bool) weatherPrediction{
	return weatherPrediction{service,saveDatabase,printPrediction}
}

func (wp *weatherPrediction)Predict() {
	ferengi := galaxy.PlanetImpl{Radius: 500, AngularSpeed: 1, ClockWise: true}
	vulcano := galaxy.PlanetImpl{Radius: 1000, AngularSpeed: 5}
	betasoide := galaxy.PlanetImpl{Radius: 2000, AngularSpeed: 3, ClockWise: true}
	sun := geometry.PointImpl{}

	var rain, drought, optimalCondition, maxRainyDay, nothing  uint
	var maxPerimeter float64
	var day uint
	var w weather.Weather
	for day = 0; day < 10*360 ; day++ {
		w.Day = day
		p1 := ferengi.AngularToPoint(day)
		p2 := vulcano.AngularToPoint(day)
		p3 := betasoide.AngularToPoint(day)

		if p1.IsCollinear(p2, p3) {
			if sun.IsCollinear(p1, p2, p3) {
				drought += 1
				w.Weather = weather.Drought
			} else {
				optimalCondition += 1
				w.Weather = weather.OptimalCondition

			}
			wp.save(w)
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
			w.Weather = weather.Rain
			wp.save(w)
			continue
		}
		w.Weather = weather.UnknowCondition
		wp.save(w)
		nothing += 1


	}
	if wp.PrintPrediction {
		fmt.Printf("Habra %v peridos de sequia \n", drought)
		fmt.Printf("Habra %v peridos de lluvia y el pico en el dia %v \n", rain, maxRainyDay)
		fmt.Printf("Habra %v peridos de condiciones optimas de clima y temperatura \n", optimalCondition)
		fmt.Printf("Habra %v peridos de condiciones desconocidas \n", nothing)
	}

}