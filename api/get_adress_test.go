package get_adress

import "testing"

func TestFetchAddressFromBrasilAPI(t *testing.T) {
	cep := "15910000"
	address, err := FetchAddressFromBrasilAPI(cep)
	if err != nil {
		t.Errorf("Erro ao buscar endereço da Brasil API: %v", err)
		return
	}

	if address == nil {
		t.Error("Endereço retornado é nulo")
		return
	}

	if address.Cep != cep {
		t.Errorf("CEP esperado: %s, CEP retornado: %s", cep, address.Cep)
	}
}

func TestFetchAddressFromViaCEP(t *testing.T) {
	cep := "15910-000"
	address, err := FetchAddressFromViaCEP(cep)
	if err != nil {
		t.Errorf("Erro ao buscar endereço da ViaCEP: %v", err)
		return
	}

	if address == nil {
		t.Error("Endereço retornado é nulo")
		return
	}

	if address.Cep != cep {
		t.Errorf("CEP esperado: %s, CEP retornado: %s", cep, address.Cep)
	}
}
