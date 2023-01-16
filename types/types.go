package types

import (
	"strings"
	"time"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        Datetime
	Assignee    User
	Duration    Duration
	Priority    string `json:"priority"`
	UserID      int    `json:"userid"`
}

type Duration struct {
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
	Months  int `json:"months"`
}

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006/01/02", strInput)
	if err != nil {
		return err
	}
	
	t.Time = newTime
	return nil
}
