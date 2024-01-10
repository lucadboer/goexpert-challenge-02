package printer_app

import (
	"fmt"
	"pos-go-luca/challenges/Multithreading/dto"
)

func PrintAddressFromBrasilAPI(address *dto.AddressBrasilAPI) {
	fmt.Println("API: Brasil API")
	fmt.Println("CEP:", address.Cep)
	fmt.Println("Estado:", address.State)
	fmt.Println("Cidade:", address.City)
}

func PrintAddressFromViaCEP(address *dto.AddressViaCEP) {
	fmt.Println("API: ViaCEP")
	fmt.Println("CEP:", address.Cep)
	fmt.Println("Logradouro:", address.Logradouro)
	fmt.Println("Complemento:", address.Complemento)
	fmt.Println("Bairro:", address.Bairro)
	fmt.Println("Localidade:", address.Localidade)
	fmt.Println("UF:", address.Uf)
}
