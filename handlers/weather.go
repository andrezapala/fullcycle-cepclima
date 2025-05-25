package handlers

import (
	"net/http"
	"regexp"

	"fullcycle-cepclima/services"

	"github.com/gin-gonic/gin"
)

func GetWeatherByCEP(c *gin.Context) {
	cep := c.Param("cep")

	if match, _ := regexp.MatchString(`^\d{8}$`, cep); !match {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	city, err := services.GetCityByCEP(cep)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	tempC, err := services.GetTemperatureByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "weather fetch failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"temp_C": tempC,
		"temp_F": tempC*1.8 + 32,
		"temp_K": tempC + 273,
	})
}
