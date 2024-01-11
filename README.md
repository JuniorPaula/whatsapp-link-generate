# Geração de Links do WhatsApp

## Pacote CLI:
Este pacote CLI permite que você gere links do WhatsApp para vários números de telefone a partir da linha de comando. Ele oferece uma maneira conveniente de gerar links do WhatsApp para contatos individuais ou em massa.

### Uso

Para usar o programa de linha de comando, siga estas etapas:

1. Navegue até o diretório `./cli/` no terminal.
2. Compile o programa usando o comando `go build main.go` ou execute diretamente com `go run main.go`.
3. Forneça os números de telefone como argumentos após o nome do programa.

#### Exemplo

```bash
# Gerar um link do WhatsApp para um único número de telefone
go run main.go 551234567890

# Gerar links do WhatsApp para vários números de telefone
go run main.go 551234567890 551234567891 551234567892
```

Os links gerados serão exibidos no console.

## Pacote API:*

Este pacote API expõe um endpoint HTTP que permite que você envie uma lista de números de telefone e uma mensagem no corpo de uma solicitação POST e receba links do WhatsApp gerados como resposta.

### Uso

Para usar a API, siga estas etapas:

1. Navegue até o diretório `./cmd/` no terminal.
2. Inicie o servidor com o comando `go run ./cmd/*`.
3. Faça uma solicitação POST para `http://localhost:8080/generate-links` com um corpo JSON contendo a lista de números de telefone e a mensagem desejada.

#### Exemplo

Suponha que você esteja usando uma ferramenta como `curl` para fazer a solicitação:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "phone_numbers": ["551234567890", "551234567891", "551234567892"],
  "message": "Olá, como posso ajudar?"
}' http://localhost:8080/generate-links
```

Ou `JSON`:
```json
{
    "phone_numbers": [
        "551234567890",
        "551234567891",
        "551234567892",
    ],
    "message": "Olá, como posso ajudar?"
}
```

A API responderá com os links gerados em formato JSON:

```json
{
  "whatsapp_links": [
    "https://api.whatsapp.com/send?phone=551234567890&text=Ol%C3%A1%2C%20como%20posso%20ajudar%3F",
    "https://api.whatsapp.com/send?phone=551234567891&text=Ol%C3%A1%2C%20como%20posso%20ajudar%3F",
    "https://api.whatsapp.com/send?phone=551234567892&text=Ol%C3%A1%2C%20como%20posso%20ajudar%3F"
  ]
}
```
