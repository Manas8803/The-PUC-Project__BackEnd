package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

// Define the data structure for the message you want to send
type WebSocketMessage struct {
	Data string `json:"data"`
}

type Payload struct {
	ConnectionId string `json:"connection_id"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data Payload
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		log.Println("Error unmarshalling payload", err)
		return events.APIGatewayProxyResponse{Body: err.Error()}, err
	}
	broadcastToConnections(data.ConnectionId, "HELLLO SOCKET")
	return events.APIGatewayProxyResponse{Body: "Invoked Successfully", StatusCode: http.StatusAccepted}, nil
}

func broadcastToConnections(connectionID string, data string) error {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("REGION")),
		Endpoint: aws.String("https://alwjqt86ih.execute-api.ap-south-1.amazonaws.com/production"),
	})
	if err != nil {
		return err
	}

	// Create a new ApiGatewayManagementApi client
	apiGatewayClient := apigatewaymanagementapi.New(sess)

	// Iterate over the connectionIDs and broadcast the data
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionID),
		Data:         []byte(data),
	}
	_, err = apiGatewayClient.PostToConnection(input)
	if err != nil {
		fmt.Printf("Error broadcasting data to connection %s: %v\n", connectionID, err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
