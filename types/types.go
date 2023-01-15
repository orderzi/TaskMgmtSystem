package types

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
	Date        string `json:"date"`
	Assigner    User   //string `json:"assigner"`
	Assignee    User   //string `json:"assignee"`
	Duration    string `json:"duration"`
	Priority    string `json:"priority"`
	UserID      int    `json:"userid"`
}
