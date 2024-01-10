package main

import (
	"fmt"
	"time"

	get_adress "github.com/lucadboer/goexpert/challenge-2/api"
	"github.com/lucadboer/goexpert/challenge-2/dto"
	printer_app "github.com/lucadboer/goexpert/challenge-2/printer"
)

func main() {
	cep := "15910000"
	brasilAPI := make(chan *dto.AddressBrasilAPI, 1)
	viaCEP := make(chan *dto.AddressViaCEP, 1)

	go func() {
		address, err := get_adress.FetchAddressFromBrasilAPI(cep)
		if err == nil {
			brasilAPI <- address
		}
	}()

	go func() {
		address, err := get_adress.FetchAddressFromViaCEP(cep)
		if err == nil {
			viaCEP <- address
		}
	}()

	select {
	case address := <-brasilAPI:
		printer_app.PrintAddressFromBrasilAPI(address)
	case address := <-viaCEP:
		printer_app.PrintAddressFromViaCEP(address)
	case <-time.After(time.Second):
		fmt.Println("Erro: Tempo de resposta excedido")
	}
}


