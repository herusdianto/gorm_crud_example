package dtos

type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations"`
}
