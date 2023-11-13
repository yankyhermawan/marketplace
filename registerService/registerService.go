package registerService

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
	"github.com/yankyhermawan/marketplace/database"
	"github.com/yankyhermawan/marketplace/interfaces"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ApiResponse struct {
	Code     int
	Response *gorm.DB
}

func RegisterUser(data *interfaces.RequestBody) ApiResponse {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitDB()

	saltStr, bool := os.LookupEnv("HASH_SALT")
	if !bool {
		log.Fatal("HASH_SALT not found")
	}
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		log.Fatal("Error converting HASH_SALT to integer:", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), salt)
	data.Password = string(hashedPassword)
	response := db.Create(data)
	return ApiResponse{Code: 200, Response: response}
}
