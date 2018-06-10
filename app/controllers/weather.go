package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/juaguz/ml-test/app/services"
)

func Clima(c *gin.Context, service services.WeatherService) (int, interface{}) {
	dia := c.Query("dia")

	d, err := strconv.Atoi(dia)

	if err != nil || d < 0 {
		return 400,gin.H{"message": "Ingrese un dia valido",}
	}

	w,err := service.Find(int64(d))

	if err != nil {
		return 500, gin.H{"message": err.Error()}
	}

	return 200,w
}