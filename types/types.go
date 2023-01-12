package types

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
}

type Users []User

type DatabaseConnection struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASS"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	DBName   string `env:"DB_NAME" envDefault:"Workout"`
}

