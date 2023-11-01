package addressapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type AddresApiCep struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

type ResultApiCep struct {
	APIName string
	Address AddresApiCep
	Error   error
}

func GetAddressFromApiCep(ctx context.Context, apiURL string, apiName string) ResultApiCep {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return ResultApiCep{APIName: apiName, Address: AddresApiCep{}, Error: err}
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	elapsed := time.Since(start)

	if err == nil && elapsed < time.Minute{
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			var address AddresApiCep
			err = json.Unmarshal(body, &address)
			if err == nil {
				return ResultApiCep{APIName: apiName, Address: address, Error: nil}
			}
		}
	}
	return ResultApiCep{APIName: apiName, Address: AddresApiCep{}, Error: err}
}
