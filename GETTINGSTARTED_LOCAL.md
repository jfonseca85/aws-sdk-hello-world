## Getting started

Instalando a versão local do AWS DynamoDB, facilitando o desenvolvimento do seu App;


#### Implantação do DynamoDB localmente em seu computador

A versão para download do Amazon DynamoDB é fornecida como um arquivo `.jar` executável. O aplicativo é executado no Windows, Linux, macOS e outras plataformas compatíveis com Java.

Siga estas etapas para configurar e executar o DynamoDB em seu computador.

**Para configurar o DynamoDB em seu computador**

* Faça download do DynamoDB gratuitamente em um dos locais a seguir.

| Região                           | Links para fazer download                                                                                | Somas de verificação                                                                             |
| ----------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| Asia Pacific (Mumbai) Region      | [.tar.gz](https://s3.ap-south-1.amazonaws.com/dynamodb-local-mumbai/dynamodb_local_latest.tar.gz)        | [.zip](https://s3.ap-south-1.amazonaws.com/dynamodb-local-mumbai/dynamodb_local_latest.zip)        |
| Asia Pacific (Singapore) Region   | [.tar.gz](https://s3.ap-southeast-1.amazonaws.com/dynamodb-local-singapore/dynamodb_local_latest.tar.gz) | [.zip](https://s3.ap-southeast-1.amazonaws.com/dynamodb-local-singapore/dynamodb_local_latest.zip) |
| Asia Pacific (Tokyo) Region       | [.tar.gz](https://s3.ap-northeast-1.amazonaws.com/dynamodb-local-tokyo/dynamodb_local_latest.tar.gz)     | [.zip](https://s3.ap-northeast-1.amazonaws.com/dynamodb-local-tokyo/dynamodb_local_latest.zip)     |
| Europe (Frankfurt) Region         | [.tar.gz](https://s3.eu-central-1.amazonaws.com/dynamodb-local-frankfurt/dynamodb_local_latest.tar.gz)   | [.zip](https://s3.eu-central-1.amazonaws.com/dynamodb-local-frankfurt/dynamodb_local_latest.zip)   |
| South America (São Paulo) Region | [.tar.gz](https://s3.sa-east-1.amazonaws.com/dynamodb-local-sao-paulo/dynamodb_local_latest.tar.gz)      | [.zip](https://s3.sa-east-1.amazonaws.com/dynamodb-local-sao-paulo/dynamodb_local_latest.zip)      |
| US West (Oregon) Region           | [.tar.gz](https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz)                | [.zip](https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.zip)                |

* Depois de fazer download do arquivo, extraia o conteúdo e copie o diretório extraído para um local de sua escolha.
* Para iniciar o DynamoDB em seu computador, abra uma janela de prompt de comando, vá para o diretório onde você extraiu o arquivo `DynamoDBLocal.jar` e insira o seguinte comando.

````
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb

````

## Go Lang: configuração das suas credenciais da AWS

O SDK for Go exige que você forneça credenciais da AWS para a sua aplicação em tempo de execução.

Veja a seguir um exemplo de um arquivo de credenciais da AWS chamado `~/.aws/credentials` em que o caractere de til (`~`) representa seu diretório inicial.

```
[default]
aws_access_key_id = `AWS access key ID goes here`
aws_secret_access_key = `Secret key goes here`
```

## Go: arquivo de configuração env.yaml
```yaml
aws:
  default_region: sa-east-1
  endpoint_url: http://localhost:8000
```

## Go: configuração da região e do endpoint da AWS

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-hello-world/configlocal"
)

...

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg.AWSClient)
```

###### Comando para criar usando aws
```sh
 aws dynamodb create-table --table-name Musicdynamodb-table-appcell  --attribute-definitions  AttributeName=ID,AttributeType=S  AttributeName=Version,AttributeType=S --key-schema AttributeName=ID,KeyType=HASH AttributeName=Version,KeyType=RANGE --billing-mode=PAY_PER_REQUEST --endpoint-url http://localhost:8000
```
###### Compile and Execute

```sh
$ go run .
Table:
hello-world-labs-table-01
```



