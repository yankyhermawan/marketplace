package database

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string
	Password string
	Name     string
}

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=marketplace port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(User{})
}
