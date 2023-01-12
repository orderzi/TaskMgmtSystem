package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orderzi/workout-service/types"
	"github.com/orderzi/workout-service/utils"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Workout page")
}

func createUser(w http.ResponseWriter, r *http.Request) (*types.User, error) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	name, err := utils.FormatName(r.FormValue("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	age, err := utils.ValidateAge(r.FormValue("age"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	weight, err := utils.RoundWeight(r.FormValue("weight"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user := &types.User{name, age, weight}
	return user, nil
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := createUser(w, r)
	if err != nil {
		return
	}
	fmt.Println(user)
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createuser", createUserHandler)
	fmt.Println("Starting server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
