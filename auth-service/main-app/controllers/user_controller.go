package controllers

import (
	"log"
	"net/http"

	"github.com/Manas8803/The-Puc-Detection/auth-service/db"
	network "github.com/Manas8803/The-Puc-Detection/auth-service/lib/net"
	"github.com/Manas8803/The-Puc-Detection/auth-service/lib/security"
	"github.com/Manas8803/The-Puc-Detection/auth-service/main-app/models"
	"github.com/Manas8803/The-Puc-Detection/auth-service/main-app/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func Login(r *gin.Context) {

	var req models.Login

	//* Checking for invalid json format
	if err := r.BindJSON(&req); err != nil {
		network.RespondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&req); validationErr != nil {
		network.RespondWithError(r, http.StatusBadRequest, "Please provide the required credentials.")
		return
	}

	//* Checking whether the user is registered
	_, userErr := db.GetUserByEmail(req.Email)
	if userErr != nil {

		log.Println(userErr)
		network.RespondWithError(r, http.StatusInternalServerError, "Internal server error : "+userErr.Error())
		return
	}

	//* Generating Token
	token, genJWTErr := security.GenerateJWT()
	if genJWTErr != nil {
		network.RespondWithError(r, http.StatusInternalServerError, "Internal Server Error : "+genJWTErr.Error())
		return
	}

	r.JSON(http.StatusOK, responses.UserResponse{Message: "success", Data: map[string]interface{}{"token": token}})
}
