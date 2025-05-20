# AWS Lambda Go API Proxy 🚀

Esta aplicação demonstra como criar uma API serverless usando Go e AWS Lambda, com o AWS Lambda Go API Proxy para integração com HTTP. ☁️

## Sobre o Projeto 📝

Este projeto utiliza o pacote `aws-lambda-go-api-proxy` para permitir que manipuladores HTTP padrão do Go sejam usados diretamente com o AWS Lambda e API Gateway, facilitando a migração de aplicativos web existentes para uma arquitetura serverless.

## Requisitos ⚙️

- Go 1.24.3 ou superior
- AWS CLI configurada
- AWS SAM (opcional, para implantação local)

## Endpoints Disponíveis 🌐

A API fornece dois endpoints:

- `/ping` - Retorna uma resposta simples "Pong" 🏓
- `/hello` - Retorna uma mensagem ASCII art com um cumprimento 👋

## Implantação no AWS Lambda 🛠️

1. Compile o binário para Linux:
    ```sh
    GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
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


