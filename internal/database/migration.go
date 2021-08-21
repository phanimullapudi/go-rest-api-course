package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/phanimullapudi/go-rest-api-course/internal/comment"
)

func MigrateDB(db *gorm.DB) error {
	fmt.Println("Testing....")
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
