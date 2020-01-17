package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")

	config = aws.Config{}
)

type response struct {
	Name      string `json:"name"`
	Country   string `json:"country"`
	IPAddress string `json:"ipAddress"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	json.Unmarshal([]byte(request.Body), response{})

	sess, err := session.NewSession(&config)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	svc := dynamodb.New(sess)
	svd := ecs.New(sess)

	item := map[string]*dynamodb.AttributeValue{
		"Id": &dynamodb.AttributeValue{S: response.Name},
	}

	svc.PutItem(&dynamodb.PutItemInput{
		Item: item,
	})

	svd

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string("")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
