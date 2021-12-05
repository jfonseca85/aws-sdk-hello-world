## Getting started

Para começar a trabalhar com o SDK, configure seu projeto para módulos Go, e recupere as dependências do SDK com `go get`

Este exemplo mostra como você pode usar o SDK v2 para fazer uma solicitação  de API  usando cliente [ Amazon DynamoDB] do SDK.

###### Iniciando o Projeto

```sh
$ mkdir ~/hellocamadazero
$ cd ~/hellocamadazero
$ go mod init hellocamadazero
```

###### Adicionando dependência do SDK

```sh
$ go get github.com/aws/aws-sdk-go-v2/aws
$ go get github.com/aws/aws-sdk-go-v2/config
$ go get github.com/aws/aws-sdk-go-v2/service/dynamodb
```

###### Escrevendo o Código

No seu editor de preferência adicione o código fonte no arquivo `main.go` sa-east-1package main

```
import (
"context"
"fmt"
"log"


"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/dynamodb"

)

func main() {
// Usando a configuração padrão do SDK's, carregando configuração adicionais
// e o valores de suas de suas credenciais  das variáveis de ambiente
cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
if err != nil {
log.Fatalf("unable to load SDK config, %v", err)
}

// Usando valor do config, criamos um client do dynamodb
svc := dynamodb.NewFromConfig(cfg)

// Construção do resquest com os parametros da configuração
resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
    Limit: aws.Int32(5),
})
if err != nil {
    log.Fatalf("failed to list tables, %v", err)
}

fmt.Println("Tables:")
for _, tableName := range resp.TableNames {
    fmt.Println(tableName)
}


}
```

###### Compile and Execute

```sh
$ go run .
Table:
camada-zero-labs-2
```
