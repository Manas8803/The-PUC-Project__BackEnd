package main

import (
	"context"
	"os"

	"github.com/Manas8803/The-Puc-Detection/auth-service/main-app/routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			Auth API
//	@version		1.0
//	@description	This is an auth api for an application.

// @BasePath	/api/v1
var ginLambda *ginadapter.GinLambda

func init() {

	prod := os.Getenv("RELEASE_MODE")
	if prod == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	godotenv.Load("../../.env")
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	//* Passing the router to all user(auth-service) routes.
	routes.UserRoute(api)

	//* Connecting to DB
	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}