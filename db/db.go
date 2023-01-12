package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/orderzi/workout-service/types"
)

func (c *types.DatabaseConnection) OpenDBSession() (*sql.DB, error) {
	if c.User == "" {
		c.User = "root"
	}
	if c.Password == "" {
		c.Password = "password"
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	if c.DBName == "" {
		c.DBName = "users"
	}

	c.User = os.Getenv("DB_USER")
	c.Password = os.Getenv("DB_PASSWORD")
	c.Host = os.Getenv("DB_HOST")
	c.Port = os.Getenv("DB_PORT")
	c.DBName = os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
