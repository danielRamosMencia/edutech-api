package constants

import (
	"log"
	"os"
	"strconv"
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
	JwtSecret   string
	JwtTime     int
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

		rawjwtTime := os.Getenv("JWT_TIME")
		jwtTime, err := strconv.Atoi(rawjwtTime)
		if err != nil {
			log.Fatal("Error converting JWT_TIME to int", err)
		}

		Envs = GlobalEnvs{
			ServerPort:  os.Getenv("SERVER_PORT"),
			Environment: os.Getenv("ENVIRONMENT"),
			DBHost:      os.Getenv("DB_HOST"),
			DBPort:      os.Getenv("DB_PORT"),
			DBUser:      os.Getenv("DB_USER"),
			DBPassword:  os.Getenv("DB_PASSWORD"),
			DBName:      os.Getenv("DB_NAME"),
			JwtSecret:   os.Getenv("JWT_SECRET"),
			JwtTime:     jwtTime,
		}
	})

	return Envs
}
