package network

import (
	"github.com/Manas8803/The-Puc-Detection/auth-service/main-app/responses"
	"github.com/gin-gonic/gin"
)

type Payload_Body struct {
	Body string `json:"body"`
}

func RespondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, responses.UserResponse{
		Message: message,
	})
}
