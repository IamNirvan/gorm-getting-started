package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IamNirvan/gorm-getting-started/db"
	"github.com/IamNirvan/gorm-getting-started/types"
	"github.com/gorilla/mux"
)

type Api struct {
	Database *db.Database
}

func main() {
	// Initialize the database
	database, databaseErr := db.NewDatabase()
	if databaseErr != nil {
		log.Fatal(databaseErr)
	}

	// Initialize Api struct
	api := &Api{
		Database: database,
	}

	// Create the user table
	entity := types.User{}
	database.InitialMigration(entity)

	// Initialize routes and handlers
	initRouter(api)
}

func initRouter(api *Api) {
	address := ":8080"

	router := mux.NewRouter()
	router.HandleFunc("/user", api.handleGetAllUsers).Methods("GET")
	router.HandleFunc("/user", api.handleCreateUser).Methods("POST").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/user/id/{id}", api.handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/user/id/{id}", api.handleUpdateUser).Methods("PATCH").HeadersRegexp("Content-Type", "application/json")

	fmt.Printf("starting web server on: %s\n", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func (a *Api) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all users")

	// Fetch all users from the database and load into slice
	var users []types.User
	a.Database.Database.Find(&users)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (a *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating users")

	// Decode the json into the user struct
	user := types.User{}
	json.NewDecoder(r.Body).Decode(&user)

	a.Database.Database.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(fmt.Sprintf("Rows affected (%d). Inserted record with id %d", result.RowsAffected, user.ID))
	json.NewEncoder(w).Encode(user)
}

func (a *Api) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user")

	// Get the path variables
	variables := mux.Vars(r)
	userId := variables["id"]

	user := types.User{}
	a.Database.Database.Where("id = ?", userId).Find(&user)
	result := a.Database.Database.Delete(&user)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Rows affected (%d)", result.RowsAffected))
}

func (a *Api) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating user")

	// Get the path variables
	variables := mux.Vars(r)
	userId := variables["id"]

	// load saved user data into user object
	user := types.User{}
	a.Database.Database.First(&user, userId)

	// override the user object with values in the json object
	json.NewDecoder(r.Body).Decode(&user)

	// save the changes in the database
	a.Database.Database.Save(&user)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
