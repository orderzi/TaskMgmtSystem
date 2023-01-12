package types

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
}

type Users []User

type DatabaseConnection struct {
	Host     string 
	User     string 
	Password string 
	Port     string   
	DBName   string 
}

