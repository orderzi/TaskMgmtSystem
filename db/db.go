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

func IsExistTask(db *sql.DB, task types.Task) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Tasks WHERE UserID = ? AND Name = ?", &task.UserID, &task.Name).Scan(&count)
	if err != nil {
		fmt.Errorf("error querying table 'Tasks'")
		return err
	}
	if count > 0 {
		return fmt.Errorf("task is already created")
	}
	return nil
}

func WriteTask(db *sql.DB, task types.Task) (sql.Result, error) {
	prep, err := db.Prepare("INSERT INTO Tasks(Name,UserID, Date, Priority, AssigneeEmail, Description, DurationMonths, DurationDays,DurationHours, DurationMinutes) VALUES (?, ? , ? , ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	writer, err := prep.Exec(task.Name, task.UserID, task.Date.Time.Format("2006/01/02"), task.Priority, task.Assignee.Email, task.Description, task.Duration.Months,
		task.Duration.Days, task.Duration.Hours, task.Duration.Minutes)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

// Writing user to DB; gets DB object and Struct User; returns string and error/nil

func WriteUser(db *sql.DB, user types.User) (sql.Result, error) {
	prep, err := db.Prepare("INSERT INTO Users(FirstName, LastName, BirthDate, Age, Email) VALUES (?, ? , ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	writer, err := prep.Exec(user.FirstName, user.LastName, user.Birthdate, user.Age, user.Email)
	if err != nil {
		return nil, err
	}
	defer prep.Close()
	return writer, nil

}

func IsExistUser(db *sql.DB, user types.User) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Users WHERE Email = ?", user.Email).Scan(&count)
	if err != nil {
		fmt.Errorf("error querying table 'Users'")
		return err
	}
	if count > 0 {
		return fmt.Errorf("user %v is already created", user.Email)
	}
	return nil
}
