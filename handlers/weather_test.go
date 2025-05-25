package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/weather/:cep", GetWeatherByCEP)
	return r
}

func TestWeatherHandler(t *testing.T) {
	os.Setenv("WEATHER_API_KEY", "dd6422816cb7468aa39141114252105")

	router := setupRouter()

	tests := []struct {
		name         string
		cep          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "CEP válido",
			cep:          "05211120",
			expectedCode: http.StatusOK,
		},
		{
			name:         "CEP inválido (formato)",
			cep:          "abc123",
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"message":"invalid zipcode"}`,
		},
		{
			name:         "CEP válido, mas não encontrado",
			cep:          "99999999",
			expectedCode: http.StatusNotFound,
			expectedBody: `{"message":"can not find zipcode"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/weather/"+tt.cep, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
