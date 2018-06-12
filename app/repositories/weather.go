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
func NewWeatherRepo(db *sql.DB) WeatherRepo {
	return &weatherRepo{db}
}
//Guarda en la base de datos el dia
// y  setea el id al objeto guardado
func (w *weatherRepo) Store(weather *weatherEntity.Weather) (error) {
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	stm := "INSERT INTO predictions (number_day, weather) VALUES ($1, $2) RETURNING id"
	result,err := tx.Exec(stm,weather.Day,weather.Weather)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	weather.Id, _ = result.LastInsertId()
	return err
}

//Busca en la base de datos el dia
func (w *weatherRepo) Find(day int64) (weatherEntity.Weather,error) {
	weather := weatherEntity.Weather{}
	stm := "SELECT * FROM predictions where number_day = $1"
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