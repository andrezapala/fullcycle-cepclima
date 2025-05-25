package main

import (
	"fullcycle-cepclima/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/weather/:cep", handlers.GetWeatherByCEP)
	r.Run(":8080")
}
