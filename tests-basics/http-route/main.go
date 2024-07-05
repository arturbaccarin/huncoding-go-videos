package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func GetCep(cep string) *ViaCEPResponse {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()

	viaCepResponse := &ViaCEPResponse{}

	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, viaCepResponse); err != nil {
		fmt.Println(err)
		return nil
	}

	return viaCepResponse
}
