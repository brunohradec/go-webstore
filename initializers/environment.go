package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

type DBEnv struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

type Env struct {
	Port string
	DB   DBEnv
}

func LoadDotenvVariables() (*Env, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	env := Env{
		Port: os.Getenv("PORT"),
		DB: DBEnv{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	return &env, nil
}
