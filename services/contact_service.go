package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/herusdianto/gorm_crud_example/dtos"
	"github.com/herusdianto/gorm_crud_example/models"
	"github.com/herusdianto/gorm_crud_example/repositories"
)

func CreateContact(contact *models.Contact, repository repositories.ContactRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	contact.ID = uuidResult.String()

	operationResult := repository.Save(contact)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}
