package repositories

import (
	"math"

	"github.com/herusdianto/gorm_crud_example/dtos"
	"github.com/herusdianto/gorm_crud_example/models"
	"github.com/jinzhu/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Save(contact *models.Contact) RepositoryResult {
	err := r.db.Save(contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: contact}
}

func (r *ContactRepository) FindAll() RepositoryResult {
	var contacts models.Contacts

	err := r.db.Find(&contacts).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contacts}
}

func (r *ContactRepository) FindOneById(id string) RepositoryResult {
	var contact models.Contact

	err := r.db.Where(&models.Contact{ID: id}).Take(&contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contact}
}

func (r *ContactRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.Contact{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ContactRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&models.Contacts{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ContactRepository) Pagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var contacts models.Contacts

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&contacts).Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = contacts

	// count all data
	errCount := r.db.Model(&models.Contact{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
