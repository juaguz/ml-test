package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juaguz/ml-test/app/services"
	"github.com/juaguz/ml-test/app/controllers"
)




func RunApi(service services.WeatherService){
	r := gin.Default()
	r.GET("/clima", func(c *gin.Context) {
		status, response := controllers.Clima(c,service)
		c.JSON(status, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}