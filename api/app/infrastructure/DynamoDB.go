package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	C *dynamodb.Client
}

func NewDynamoDBClient() *DynamoDBClient {
	access_key_id := os.Getenv("DYNAMO_ACCESS_KEY_ID")
	secret_access_key := os.Getenv("DYNAMO_SECRET_ACCESS_KEY")
	dynamodb_region := os.Getenv("DYNAMO_REGION")
	// db_port := os.Getenv("DYNAMO_DB_PORT")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(dynamodb_region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://dynamodb-local:8000"}, nil
			},
		)),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     access_key_id,
				SecretAccessKey: secret_access_key,
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)
	tables, err := client.ListTables(
		context.TODO(), &dynamodb.ListTablesInput{},
	)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(tables.TableNames)
	}

	return &DynamoDBClient{C: client}
}
