package db

import (
	"github.com/khoahotran/ticket-booking-project-v0/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}