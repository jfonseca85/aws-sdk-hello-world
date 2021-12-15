package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	tableName := "hello-world-labs-table-01"
	attributeName := "id-labs-table-01"

	var keytable []types.KeySchemaElement
	var schemaElement = types.KeySchemaElement{
		AttributeName: &attributeName,
		KeyType:       types.KeyTypeHash,
	}
	keytable = append(keytable, schemaElement)

	var attributeDefinitionList []types.AttributeDefinition
	var attributeDefinition = types.AttributeDefinition{
		AttributeName: &attributeName,
		AttributeType: "S",
	}
	attributeDefinitionList = append(attributeDefinitionList, attributeDefinition)

	createTableOutput := dynamodb.CreateTableInput{

		AttributeDefinitions: attributeDefinitionList,
		TableName:            &tableName,
		KeySchema:            keytable,
		BillingMode:          types.BillingModePayPerRequest,
	}

	_, err = svc.CreateTable(context.TODO(), &createTableOutput)

	// Build the request with its input parameters
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
