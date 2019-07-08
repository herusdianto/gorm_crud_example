package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herusdianto/gorm_crud_example/dtos"
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

	route.GET("/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllContacts(*contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindOneContactById(id, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var contact models.Contact

		err := context.ShouldBindJSON(&contact)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateContactById(id, &contact, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteOneContactById(id, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		if len(multiID.Ids) == 0 {
			response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.DeleteContactByIds(&multiID, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequest(context)

		response := services.Pagination(*contactRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	return route
}
