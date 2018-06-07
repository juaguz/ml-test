package app

import "github.com/juaguz/ml-test/app/entities/galaxy"

func WeatherPrediction(day uint) {
	ferengi := galaxy.PlanetImpl{500,1, true}
	vulcano := galaxy.PlanetImpl{1000,5, false}
	betasoide := galaxy.PlanetImpl{2000,3, true}

	for day = 0; day <= 10*364 ; day++ {
		p1 := ferengi.AngularToPoint(day)
		p2 := vulcano.AngularToPoint(day)
		p3 := betasoide.AngularToPoint(day)

		isInline := p1.IsInline(p2,p3)


	}

}
