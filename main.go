package main

import ("flag"
	"database/sql"
	"github.com/juaguz/ml-test/app/config"
	"fmt"
	"github.com/juaguz/ml-test/app/repositories"
	"github.com/juaguz/ml-test/app/services"

	"github.com/juaguz/ml-test/app/api"
	"github.com/juaguz/ml-test/app/utils"
)


func main(){
	var weatherService services.WeatherService

	printPrediction := flag.Bool("print", false, "imprime las predicciones")

	fillDatabase    := flag.Bool("fill", false, "llena la base de datos")

	startApi		:= flag.Bool("api", false, "llena la base de datos")

	flag.Parse()

	c:= config.NewConfig("Weather")


	if *fillDatabase || *startApi {
		dbConn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.DbName)

		db, err := sql.Open(c.Driver, dbConn)
		defer db.Close()
		utils.CreateTable(db)
		if err != nil {
			panic(err)
		}
		weatherRepo := repositories.NewWeatherRepo(db)
		weatherService = services.NewWeatherService(weatherRepo)
	}



	p := services.NewWeatherPrediction(weatherService,*fillDatabase, *printPrediction)
	p.Predict()
	if *startApi {
		api.RunApi(weatherService)
	}

}


