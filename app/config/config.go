package config

import "github.com/juaguz/config"

type Config struct {
    Database string
	Driver string
}



func NewConfig(prefix string) Config {
	return initVipper(prefix)
}

func initVipper(prefix string) Config {
	c := Config{}
	config.AddConfigPath("/etc/config")
	config.SetConfig(prefix, &c, "config")
	return c
}


//WEATHER_HOST=127.0.0.1 WEATHER_PORT=5432 WEATHER_DBNAME=weather WEATHER_PASSOWRD=ml go run  main.go -print=true
