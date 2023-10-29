package apiresult

import (
	"fmt"

	"github.com/lgustavopalmieri/challenge-multithreading/addressapi"
)

func PrintAddressDetails(result addressapi.Result) {
	if result.Error == nil {
		fmt.Printf("API mais rápida: %s\n", result.APIName)
		fmt.Printf("CEP: %s\n", result.Address.Cep)
		fmt.Printf("Logradouro: %s\n", result.Address.Logradouro)
		fmt.Printf("Bairro: %s\n", result.Address.Bairro)
		fmt.Printf("Cidade: %s\n", result.Address.Localidade)
		fmt.Printf("UF: %s\n", result.Address.Uf)
	} else {
		fmt.Printf("Erro ao obter dados do endereço: %v\n", result.Error)
	}
}
