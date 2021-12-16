package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-hello-world/configlocal"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := dynamodb.NewFromConfig(cfg.AWSClient)

	/* Subir local sem env.yaml
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)
	// Using the Config value, create the DynamoDB client
	*/

	tableName := "dynamodb-table-appcell"

	attributeNameID := "ID"
	var schemaElementID = types.KeySchemaElement{
		AttributeName: &attributeNameID,
		KeyType:       types.KeyTypeHash,
	}

	attributeNameVersion := "Version"
	var schemaElementVersion = types.KeySchemaElement{
		AttributeName: &attributeNameVersion,
		KeyType:       types.KeyTypeRange,
	}

	attributeNameStatus := "Status"
	/*
		var schemaElementStatus = types.KeySchemaElement{
			AttributeName: &attributeNameStatus,
			KeyType:       types.KeyTypeRange,
		}*/

	var keytable []types.KeySchemaElement
	keytable = append(keytable, schemaElementID, schemaElementVersion)

	var attributeDefinitionID = types.AttributeDefinition{
		AttributeName: &attributeNameID,
		AttributeType: "S",
	}
	var attributeDefinitionVersion = types.AttributeDefinition{
		AttributeName: &attributeNameVersion,
		AttributeType: "S",
	}

	var attributeDefinitionStatus = types.AttributeDefinition{
		AttributeName: &attributeNameStatus,
		AttributeType: "S",
	}
	var attributeDefinitionList []types.AttributeDefinition
	attributeDefinitionList = append(attributeDefinitionList, attributeDefinitionID, attributeDefinitionVersion, attributeDefinitionStatus)

	createTableOutput := dynamodb.CreateTableInput{

		AttributeDefinitions: attributeDefinitionList,
		TableName:            &tableName,
		KeySchema:            keytable,
		BillingMode:          types.BillingModePayPerRequest,
	}

	svc.CreateTable(context.TODO(), &createTableOutput)

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
