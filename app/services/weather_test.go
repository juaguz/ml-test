package services

import (
	"testing"
	"github.com/juaguz/ml-test/app/repositories/mocks"
	"github.com/juaguz/ml-test/app/models/weather"
)

func TestWeatherServiceStore(t *testing.T) {
	weatherRepo := new(mocks.WeatherRepo)
	weatherService := NewWeatherService(weatherRepo)
	weatherEntity := weather.Weather{Day:1,Weather:weather.OptimalCondition}
	weatherRepo.On("Store",&weatherEntity).Return(nil)
	err := weatherService.Store(&weatherEntity)
	if err != nil {
		t.Errorf("the returned err should be %v not : %v", nil,err)
	}
}

func TestWeatherServiceStoreValidation(t *testing.T) {
	weatherRepo := new(mocks.WeatherRepo)
	weatherService := NewWeatherService(weatherRepo)
	weatherEntity := weather.Weather{}
	err := weatherService.Store(&weatherEntity)
	if err.Error()!= "Clima invalido" {
		t.Fatalf("error should be Clima invalido, not %v", err.Error())
	}
}
