package db

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/orderzi/workout-service/types"
)

// var host string = os.Getenv("DB_HOST")
// var User string = os.Getenv("DB_USER")
// var Password string = os.Getenv("DB_PASSWORD")
// var Port int = os.Getenv("DB_PORT")
// var DBName string = os.Getenv("DB_NAME")



func OpenDbSession(dbconfig types.DatabaseConnection){
	sql , err := mysql.
}