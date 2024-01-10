package get_adress

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucadboer/goexpert/challenge-2/dto"
)

func FetchAddressFromBrasilAPI(cep string) (*dto.AddressBrasilAPI, error) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var address dto.AddressBrasilAPI
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func FetchAddressFromViaCEP(cep string) (*dto.AddressViaCEP, error) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var address dto.AddressViaCEP
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
