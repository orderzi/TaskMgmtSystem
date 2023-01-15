package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/orderzi/workout-service/db"
	"github.com/orderzi/workout-service/types"
	"github.com/orderzi/workout-service/utils"
)

var db_user string = os.Getenv("DB_USER")
var db_password string = os.Getenv("DB_PASSWORD")
var db_host string = os.Getenv("DB_HOST")
var db_port string = os.Getenv("DB_PORT")
var db_name string = os.Getenv("DB_NAME")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to MyHome")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	firstname, err := utils.FormatName(r.FormValue("firstname"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastname, err := utils.FormatName(r.FormValue("lastname"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	birthdate, err := utils.ValidateDate(r.FormValue("birthdate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &types.User{
		FirstName: firstname,
		LastName:  lastname,
		Birthdate: birthdate,
	}
	err = utils.SetAge(user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	db_params := db.DatabaseConnection{
		Host:     db_host,
		User:     db_user,
		Password: db_password,
		Port:     db_port,
		DBName:   db_name,
	}
	connection, err := db_params.OpenDBSession()
	if err != nil {
		http.Error(w, "Cannot establish DB connection", http.StatusBadRequest)
		return
	}

	_, err = db.WriteUser(connection, *user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(writer)

}


func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createuser", createUser)
	fmt.Println("Starting server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
