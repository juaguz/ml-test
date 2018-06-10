package repositories

import (
	weatherEntity "github.com/juaguz/ml-test/app/models/weather"
	"database/sql"
	"log"
)

type weatherRepo struct {
	DB *sql.DB
}

//Inicializador del repo
func NewWeatherRepo(session *sql.DB) WeatherRepo {
	return &weatherRepo{session}
}
//Guarda en la base de datos el dia
// y  setea el id al objeto guardado
func (w *weatherRepo) Store(weather *weatherEntity.Weather) (error) {
	stm := "INSERT INTO predictions (number_day, weather) VALUES ($1, $2) RETURNING id"
	err := w.DB.QueryRow(stm,weather.Day,weather.Weather).Scan(&weather.Id)
	return err
}

//Busca en la base de datos el dia
func (w *weatherRepo) Find(day int64) (weatherEntity.Weather,error) {
	weather := weatherEntity.Weather{}
	stm := "SELECT id, number_day, weather FROM predictions where number_day = $1"
	rows, err := w.DB.Query(stm, day)
	if err != nil {
		log.Println(err)
		return weather, err
	}
	for rows.Next() {
		rows.Scan(&weather.Id, &weather.Day, &weather.Weather)
	}
	return weather,nil
}