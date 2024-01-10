# Multithreading Challenge

Este é um desafio de programação em Go que consiste em buscar informações de endereço de duas APIs diferentes simultaneamente e imprimir os detalhes do endereço da API que responder mais rapidamente.

## Funcionalidades

- Consulta de endereço a partir de dois serviços: BrasilAPI e ViaCEP.
- Impressão dos detalhes do endereço retornado pela API mais rápida.

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [HTTP Client](https://pkg.go.dev/net/http)
- [JSON Encoding/Decoding](https://pkg.go.dev/encoding/json)
- [Concorrência com Goroutines](https://tour.golang.org/concurrency/1)

## Como Rodar

1. **Pré-requisitos**:
    - Certifique-se de ter o Go instalado em sua máquina. Caso contrário, você pode baixá-lo [aqui](https://golang.org/dl/).

2. **Clonar o Repositório**:
    ```bash
    git clone https://github.com/seu_usuario/pos-go-luca.git
    cd pos-go-luca/challenges/Multithreading
    ```

3. **Executar a Aplicação**:
    ```bash
    go run main.go
    ```

## Estrutura do Projeto

- `api/`: Contém as funções para buscar endereços de diferentes APIs.
- `printer_app/`: Funções para imprimir os detalhes do endereço.
- `dto/`: Definições de tipos de dados relacionados a endereços.

## Contribuição

Se você encontrar algum problema ou tiver alguma sugestão para melhorar este projeto, sinta-se à vontade para abrir uma [issue](https://github.com/seu_usuario/pos-go-luca/issues) ou enviar um pull request.
