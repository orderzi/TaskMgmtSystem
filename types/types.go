package types

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
}

type Task struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Assigner string `json:"asigner"`
	Assignee string `json:"asignee"`
}
