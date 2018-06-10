package repositories

import "github.com/juaguz/ml-test/app/models/weather"

type WeatherRepo interface {
	Store(weather *weather.Weather) error
	Find(day int64) (weather.Weather, error)
}