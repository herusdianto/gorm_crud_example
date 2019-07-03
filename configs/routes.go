package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herusdianto/gorm_crud_example/helpers"
	"github.com/herusdianto/gorm_crud_example/models"
	"github.com/herusdianto/gorm_crud_example/repositories"
	"github.com/herusdianto/gorm_crud_example/services"
)

func SetupRoutes(contactRepository *repositories.ContactRepository) *gin.Engine {
	route := gin.Default()

	// create route /create endpoint
	route.POST("/create", func(context *gin.Context) {
		// initialization contact model
		var contact models.Contact

		// validate json
		err := context.ShouldBindJSON(&contact)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateContact(&contact, *contactRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	return route
}
