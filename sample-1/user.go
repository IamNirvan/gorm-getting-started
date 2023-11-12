package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Email     string `json:"email"`
// }

// var DB *gorm.DB
// var DBErr error

// func initialMigration() {
// 	dsn := "user=postgres password=shalin123 host=localhost port=5432 dbname=gorm_getting_started sslmode=disable"
// 	DB, DBErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if DBErr != nil {
// 		log.Fatal(DBErr)
// 	}

// 	DB.AutoMigrate(&User{})
// }

// func handleGetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User

// 	// Decode the json to the user struct
// 	json.NewDecoder(r.Body).Decode(&user)

// 	// Save to the database
// 	DB.Create(&user)

// 	// Encode the saved user object to JSON to send as a response
// 	json.NewEncoder(w).Encode(user)
// }

// func handleCreateUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("something")
// }

// func handleUpdateUsers(w http.ResponseWriter, r *http.Request) {

// }

// func handleDeleteUsers(w http.ResponseWriter, r *http.Request) {

// }
