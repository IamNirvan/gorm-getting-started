package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/user", handleGetUsers).Methods("GET")
	router.HandleFunc("/user/id/{id}", handleGetUser).Methods("GET")
	router.HandleFunc("/user", handleCreateUsers).Methods("POST")
	router.HandleFunc("/user/id/{id}", handleUpdateUsers).Methods("PUT")
	router.HandleFunc("/user/id/{id}", handleDeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	initialMigration()
	initRouter()
}

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var DBErr error

func initialMigration() {
	dsn := "user=postgres password=shalin123 host=localhost port=5432 dbname=gorm_getting_started sslmode=disable"
	DB, DBErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if DBErr != nil {
		log.Fatal(DBErr)
	}

	DB.AutoMigrate(&User{})
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userList []User
	DB.Find(&userList)

	json.NewEncoder(w).Encode(userList)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVariables := mux.Vars(r)
	var user User
	DB.First(&user, pathVariables["id"])
	json.NewEncoder(w).Encode(user)
}

func handleCreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	// Decode the json to the user struct
	json.NewDecoder(r.Body).Decode(&user)

	// Save to the database
	DB.Create(&user)

	// Encode the saved user object to JSON to send as a response
	json.NewEncoder(w).Encode(user)
}

func handleUpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVariables := mux.Vars(r)
	var user User
	DB.First(&user, pathVariables["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func handleDeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathVariables := mux.Vars(r)
	var user User
	DB.Delete(&user, pathVariables["id"])
	json.NewEncoder(w).Encode("user with ID " + pathVariables["id"] + " was deleted")
}
