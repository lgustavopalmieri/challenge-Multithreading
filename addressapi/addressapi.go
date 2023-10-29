package addressapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type Result struct {
	APIName string
	Address Address
	Error   error
}

func GetAddressFromAPI(ctx context.Context, apiURL string, apiName string) Result {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return Result{APIName: apiName, Address: Address{}, Error: err}
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	elapsed := time.Since(start)

	if err == nil && elapsed < time.Second {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			var address Address
			err = json.Unmarshal(body, &address)
			if err == nil {
				return Result{APIName: apiName, Address: address, Error: nil}
			}
		}
	}

	return Result{APIName: apiName, Address: Address{}, Error: err}
}
