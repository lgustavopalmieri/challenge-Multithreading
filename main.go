package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lgustavopalmieri/challenge-multithreading/addressapi"
	"github.com/lgustavopalmieri/challenge-multithreading/apiresult"
	"github.com/lgustavopalmieri/challenge-multithreading/userinput"
)

func main() {
	for {
		cep, err := userinput.GetUserInput()
		if err != nil {
			fmt.Printf("Erro ao obter entrada do usuário: %v\n", err)
			return
		}
		apicep := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
		viacep := "http://viacep.com.br/ws/" + cep + "/json/"

		ch := make(chan interface{})
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		go func() {
			ch <- addressapi.GetAddressFromApiCep(ctx, apicep, "apicep")
		}()

		go func() {
			ch <- addressapi.GetAddressFromViaCep(ctx, viacep, "viacep")
		}()

		select {
		case result := <-ch:
			if result, ok := result.(addressapi.ResultApiCep); ok {
				apiresult.PrintApiCepResult(result)
			}
			if result, ok := result.(addressapi.ResultViaCep); ok {
				apiresult.PrintViaCepResult(result)
			}
		case <-ctx.Done():
			fmt.Println("Timeout ao buscar dados do endereço.")
		}
	}
}
