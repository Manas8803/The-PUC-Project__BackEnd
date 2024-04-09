package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleConnection(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Disconnecting Connection ID : ", event.RequestContext.ConnectionID)

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func main() {
	lambda.Start(handleConnection)
}
