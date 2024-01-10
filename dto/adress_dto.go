package dto

type AddressBrasilAPI struct {
	Cep      string `json:"cep"`
	State    string `json:"state"`
	City     string `json:"city"`
}

type AddressViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}
