package services

import (
	"github.com/juaguz/ml-test/app/repositories"
	weatherEntity "github.com/juaguz/ml-test/app/models/weather"
	"github.com/pkg/errors"
)

type weatherService struct {
	repositories.WeatherRepo
}

func NewWeatherService(weatherRepo repositories.WeatherRepo) WeatherService {
	return &weatherService{weatherRepo}
}

func (w *weatherService) Store(weather *weatherEntity.Weather) (error) {
	//Valida el dia
	if weather.Day < 0 {
		return errors.New("Dia invalido")
	}
	//Valido el clima
	if weather.Weather == "" {
		return errors.New("Clima invalido")
	}
	return w.WeatherRepo.Store(weather)
}

func (w *weatherService) Find(day int64) (weatherEntity.Weather,error) {
	var weather weatherEntity.Weather
	if day < 0 {
		return weather,errors.New("Dia invalido")
	}
	weather, err :=  w.WeatherRepo.Find(day)
	if err !=nil {
		return weather, err
	}
	if weather.Id <= 0 {
		return weather, errors.New("Dia no encontrado")
	}
	return weather,nil

}
