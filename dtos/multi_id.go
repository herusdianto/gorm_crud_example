package dtos

type MultiID struct {
	Ids []string `json:"ids" binding:"required"`
}
