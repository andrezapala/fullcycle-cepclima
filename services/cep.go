package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

func GetCityByCEP(cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", errors.New("city not found")
	}

	return data.Localidade, nil
}
