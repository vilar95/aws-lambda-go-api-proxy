# AWS Lambda Go API Proxy 🚀

Esta aplicação demonstra como criar uma API serverless usando Go e AWS Lambda, com o AWS Lambda Go API Proxy para integração com HTTP. ☁️

## Sobre o Projeto 📝

Este projeto utiliza o pacote `aws-lambda-go-api-proxy` para permitir que manipuladores HTTP padrão do Go sejam usados diretamente com o AWS Lambda e API Gateway, facilitando a migração de aplicativos web existentes para uma arquitetura serverless.

## Requisitos ⚙️

- Go 1.24.3 ou superior
- AWS CLI configurada
- AWS SAM (opcional, para implantação local)

## Configuração dos Pacotes 📦

O `go mod` é o sistema de gerenciamento de dependências do Go. Ele permite que você defina, baixe e controle as versões dos pacotes usados no seu projeto. 🔄

### Como funciona a configuração dos pacotes com `go mod`: 🛠️

1. **Inicialização do Módulo** 🚀  
  No terminal, execute:
  ```sh
  go mod init nome-do-modulo
  ```
  Isso cria um arquivo `go.mod` na raiz do projeto, que registra o nome do módulo e as dependências.

2. **Adicionando Dependências** ➕  
  Ao importar um pacote externo no seu código e rodar `go build` ou `go run`, o Go adiciona automaticamente a dependência ao `go.mod` e baixa o pacote.  
  Exemplo:
  ```go
  import "github.com/aws/aws-lambda-go"
  ```
  Depois de rodar:
  ```sh
  go build
  ```
  O `go.mod` será atualizado com a dependência e sua versão.

3. **Arquivo `go.sum`** 📝  
  O Go cria/atualiza o arquivo `go.sum` para garantir a integridade das dependências baixadas. 🔒

4. **Removendo Dependências Não Utilizadas** 🧹  
  Execute:
  ```sh
  go mod tidy
  ```
  Isso remove dependências não utilizadas e adiciona as que faltam.

### Exemplo de `go.mod`: 📄

```go
require (
   github.com/aws/aws-lambda-go v1.48.0
   github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
)
```

## Endpoints Disponíveis 🌐

A API fornece dois endpoints:

- `/ping` - Retorna uma resposta simples "Pong" 🏓
- `/hello` - Retorna uma mensagem ASCII art com um cumprimento 👋

## Implantação no AWS Lambda 🛠️

1. Compile o binário para Linux:
    ```sh
    env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w -X 'main.env=$(ENV)'" -o bin/bootstrap .
    ```

2. Crie um arquivo ZIP contendo o executável:
    ```sh
    zip bootstrap.zip bootstrap
    ```

3. Faça upload do arquivo `bootstrap.zip` como um novo Lambda Function no console da AWS, escolhendo o runtime "Go".

4. Configure as permissões e variáveis de ambiente conforme necessário.

5. (Opcional) Utilize o AWS SAM ou CLI para automatizar o deploy.

## Como Executar Localmente 🖥️

1. Compile isso:
    ```
    curl -X GET "https://"Endereço da API Gateway".amazonaws.com/"Seu Ambiente configurado"/ping" \
  -H "x-api-key: "Sua chave da API" "
    ```

    ```
    curl -X GET "https://"Endereço da API Gateway".amazonaws.com/"Seu Ambiente configurado"/hello" \
  -H "x-api-key: "Sua chave da API" "
    ```

## Tecnologias Utilizadas 🧰

- [Go](https://golang.org/) 🦫
- [AWS Lambda](https://aws.amazon.com/lambda/) 🟡
- [aws-lambda-go](https://github.com/aws/aws-lambda-go) 📦
- [aws-lambda-go-api-proxy](https://github.com/awslabs/aws-lambda-go-api-proxy) 🔗


