package models

import "time"

type Contact struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Phone     string `gorm:"type:varchar(100);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"phone" binding:"required"`
	Address   string `gorm:"type:text" json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contacts []Contact
