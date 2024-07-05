package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetCep(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cep := "01010101"

	httpmock.RegisterResponder("GET", fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep),
		func(req *http.Request) (*http.Response, error) {
			ViaCEPResponse := ViaCEPResponse{
				CEP:         "021549874",
				Logradouro:  "Rua teste",
				Complemento: "",
				Bairro:      "Centro",
				Localidade:  "São Paulo",
				UF:          "SP",
				IBGE:        "3550308",
				GIA:         "",
				DDD:         "11",
				SIAFI:       "7107",
			}

			resp, _ := httpmock.NewJsonResponse(200, ViaCEPResponse)
			return resp, nil
		})

	response := GetCep(cep)

	if response == nil {
		t.Fatalf("error trying http mock")
	}
}
