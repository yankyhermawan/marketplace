package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"strconv"

	"github.com/joho/godotenv"
	"github.com/yankyhermawan/marketplace/database"
	"github.com/yankyhermawan/marketplace/interfaces"
	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

type ApiResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}

func GetAllUsers() ApiResponse {
	db := database.InitDB()

	var user []database.User
	response := db.Find(&user)
	if response.Error != nil {
		return ApiResponse{
			Code:     http.StatusInternalServerError,
			Response: "Query failed",
		}
	}
	if response.RowsAffected == 0 {
		return ApiResponse{
			Code:     http.StatusNotFound,
			Response: "No records found",
		}
	}
	fmt.Println(response)
	return ApiResponse{Code: 200, Response: response}
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
	user := database.User{Name: data.Name, Password: data.Password, Username: data.Username}
	response := db.Create(&user)
	if response.Error != nil {
		return ApiResponse{Code: 500, Response: "Error creating user"}
	}
	return ApiResponse{Code: 200, Response: response}
}
