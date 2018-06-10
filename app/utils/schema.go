package utils

import "database/sql"

func CreateTable(db *sql.DB){
	create := "CREATE TABLE IF NOT EXISTS public.predictions ( id bigserial, number_day bigint, weather character(25), PRIMARY KEY (id)) WITH (OIDS = FALSE);"
	stm, err := db.Prepare(create)
	if err != nil {
		panic(err)
	}
	_, err = stm.Exec()

	if err != nil {
		panic(err)
	}
}
