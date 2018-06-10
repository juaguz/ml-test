package main

import (
	"database/sql"
	"flag"
	"github.com/juaguz/ml-test/app/api"
	"github.com/juaguz/ml-test/app/config"
	"github.com/juaguz/ml-test/app/repositories"
	"github.com/juaguz/ml-test/app/services"
	"github.com/juaguz/ml-test/app/utils"
	_ "github.com/lib/pq"
)

func main() {
	var weatherService services.WeatherService

	printPrediction := flag.Bool("print", false, "imprime las predicciones")

	fillDatabase := flag.Bool("fill", false, "llena la base de datos")

	startApi := flag.Bool("api", false, "llena la base de datos")

	flag.Parse()

	c := config.NewConfig("Weather")

	if *fillDatabase || *startApi {


		db, err := sql.Open(c.Driver, c.Database)
		defer db.Close()

		if err != nil {
			panic(err)
		}
		utils.CreateTable(db)
		weatherRepo := repositories.NewWeatherRepo(db)
		weatherService = services.NewWeatherService(weatherRepo)
	}

	p := services.NewWeatherPrediction(weatherService, *fillDatabase, *printPrediction)
	p.Predict()
	if *startApi {
		api.RunApi(weatherService)
	}

}
