package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Digite o CEP ou aperte CTRL+C para encerrar: ")
	cepInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	cep := strings.TrimSpace(cepInput)
	return cep, nil
}
