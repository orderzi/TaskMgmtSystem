package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/orderzi/workout-service/types"
)

type DatabaseConnection struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

// Open a DB session with ENV_VARS with fallback to defaults; returns DB object + error\nil

func (c *DatabaseConnection) OpenDBSession() (*sql.DB, error) {
	if c.User == "" {
		c.User = "workout_user"
	}
	if c.Password == "" {
		c.Password = "mypassword"
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	if c.DBName == "" {
		c.DBName = "workout"
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Writing user to DB; gets DB object and Struct User; returns string and error/nil

func WriteUser(db *sql.DB, user types.User) (sql.Result, error) {
	prep, err := db.Prepare("INSERT INTO Users(FirstName, LastName, BirthDate, Age) VALUES (?, ? , ?, ?) ON DUPLICATE KEY UPDATE FirstName=?, LastName=?, BirthDate=?, Age=?")
	if err != nil {
		return nil, err
	}
	writer, err := prep.Exec(user.FirstName, user.LastName, user.Birthdate, user.Age, user.FirstName, user.LastName, user.Birthdate, user.Age)
	if err != nil {
		return nil, err
	}
	defer prep.Close()
	return writer, nil

}
