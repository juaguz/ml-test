package services

import "github.com/juaguz/ml-test/app/models/weather"

type WeatherService interface {
	Store(weather *weather.Weather) error
	Find(day int64) (weather.Weather, error)
}
