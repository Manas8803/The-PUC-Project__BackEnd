package routes

import (
	"github.com/Manas8803/The-Puc-Detection/auth-service/main-app/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup) {
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
}
