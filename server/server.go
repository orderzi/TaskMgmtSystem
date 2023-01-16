package server

import (
	"encoding/json"
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

	email, err := utils.ValidateEmail(r.FormValue("email"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &types.User{
		FirstName: firstname,
		LastName:  lastname,
		Birthdate: birthdate,
		Email:     email,
	}
	err = utils.SetAge(user)
	if err != nil {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	is_user_exist := db.IsExistUser(connection, *user)
	if is_user_exist != nil {
		http.Error(w, is_user_exist.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.WriteUser(connection, *user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	task.Name = r.FormValue("name")
	if r.FormValue("name") == "" {
		http.Error(w, "Task name should not be blank", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Error decoding request body"+err.Error(), http.StatusBadRequest)
		return
	}
	db_params := db.DatabaseConnection{
		Host:     db_host,
		User:     db_user,
		Password: db_password,
		Port:     db_port,
		DBName:   db_name,
	}
	connection, err := db_params.OpenDBSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(task)
	is_task_exist := db.IsExistTask(connection, task)
	if is_task_exist != nil {
		http.Error(w, "task is already created", http.StatusBadRequest)
		return
	}

	_, err = db.WriteTask(connection, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createuser", createUser)
	myRouter.HandleFunc("/createtask", createTask)

	fmt.Println("Starting server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
