package addressapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type AddressViaCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type ResultViaCep struct {
	APIName string
	Address AddressViaCep
	Error   error
}

func GetAddressFromViaCep(ctx context.Context, apiURL string, apiName string) ResultViaCep {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return ResultViaCep{APIName: apiName, Address: AddressViaCep{}, Error: err}
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	elapsed := time.Since(start)

	if err == nil && elapsed < time.Minute {
		defer resp.Body.Close()

		var address AddressViaCep
		err = json.NewDecoder(resp.Body).Decode(&address)
		if err == nil {
			return ResultViaCep{APIName: apiName, Address: address, Error: nil}
		}
	}

	return ResultViaCep{APIName: apiName, Address: AddressViaCep{}, Error: err}
}
