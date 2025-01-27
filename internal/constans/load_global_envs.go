package constans

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type GlobalEnvs struct {
	ServerPort  string
	Environment string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

var (
	Envs GlobalEnvs
	once sync.Once
	err  error
)

func LoadGlobalEnvs() GlobalEnvs {
	once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}

		Envs = GlobalEnvs{
			ServerPort:  os.Getenv("SERVER_PORT"),
			Environment: os.Getenv("ENVIRONMENT"),
			DBHost:      os.Getenv("DB_HOST"),
			DBPort:      os.Getenv("DB_PORT"),
			DBUser:      os.Getenv("DB_USER"),
			DBPassword:  os.Getenv("DB_PASSWORD"),
			DBName:      os.Getenv("DB_NAME"),
		}
	})

	return Envs
}
