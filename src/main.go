package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	errNameNotProvided = errors.New("no name provided in HTTP body")
)

func init() {
	fmt.Println("Init")
}

// Handler is the primary Lambda handler for the API
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errNameNotProvided
	}
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	fmt.Println("Main")
	lambda.Start(Handler)
}
