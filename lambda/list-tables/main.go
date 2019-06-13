package main

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// List tables
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint := os.Getenv("DYNAMO_ENDPOINT")
	region := os.Getenv("DYNAMO_REGION")
	fmt.Println(endpoint, region)

	svc := dynamodb.New(session.New(), aws.NewConfig().
		WithEndpoint(endpoint).
		WithRegion(region))

	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := struct {
		TableNames []*string `json:"table_names"`
	}{
		result.TableNames,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
