package models

import "time"

type Contact struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);NOT NULL"`
	Email     string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(100);NOT NULL;UNIQUE;UNIQUE_INDEX"`
	Address   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contacts []Contact
