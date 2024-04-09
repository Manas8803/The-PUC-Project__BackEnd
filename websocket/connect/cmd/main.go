package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleConnection(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Connection ID : ", event.RequestContext.ConnectionID)

	// You can add additional logic here to handle the connection event

	// Return a successful response with HTTP status code 200
	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func main() {
	lambda.Start(handleConnection)
}
