package repositories

import (
	"testing"
	"database/sql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/juaguz/ml-test/app/models/weather"
)

func mockDb(t *testing.T) (*sql.DB, sqlmock.Sqlmock){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db,mock
}




func TestWeatherRepoStore(t *testing.T) {
	db, mock := mockDb(t)
	defer db.Close()
	weatherModel := weather.Weather{Weather:weather.OptimalCondition,Day:10}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO predictions").WithArgs(weatherModel.Day,weatherModel.Weather).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	weatherRepo := NewWeatherRepo(db)
	weatherRepo.Store(&weatherModel)
	if weatherModel.Id != 1 {
		t.Errorf("the returned id should be %v not : %v", 1,weatherModel.Id)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
	 t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestWeatherRepoFind(t *testing.T) {
	db, mock := mockDb(t)
	defer db.Close()
	mock.ExpectQuery("SELECT (.+) FROM predictions (.+) ").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id","weather","day"}).AddRow(1,weather.OptimalCondition,10))
	weatherRepo := NewWeatherRepo(db)
	weatherModel, err := weatherRepo.Find(1)
	if err != nil {
		t.Errorf("the returned err should be %v not : %v", nil,err)
	}
	if weatherModel.Id != 1 {
		t.Errorf("the returned id should be %v not : %v", 1,weatherModel.Id)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}