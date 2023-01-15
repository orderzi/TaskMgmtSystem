package types

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

type Task struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Assigner string `json:"assigner"`
	Assignee string `json:"assignee"`
}
