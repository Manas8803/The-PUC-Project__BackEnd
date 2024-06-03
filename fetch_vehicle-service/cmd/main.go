package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Manas8803/The-PUC-Project__BackEnd/fetch_vehicle-service/pkg/models/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Officename string `json:"office_name"`
}

type Response struct {
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data Payload
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		log.Println("Error in unmarshalling data : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	vehicles, err := service.FetchVehicles(data.Officename)
	if err != nil {
		log.Println("Error in fetching vehicles : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	payload, err := json.Marshal(vehicles)
	if err != nil {
		log.Println("Error in marshalling vehicles : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	resp := Response{
		Message: "Successfully fetched vehicles from DB",
		Data:    payload,
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error in marshalling response : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{Body: string(respBytes), StatusCode: http.StatusOK}, nil
}
func main() {
	lambda.Start(Handler)
}
