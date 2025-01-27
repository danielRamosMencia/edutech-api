package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/danielRamosMencia/edutech-api/internal/constans"

	_ "github.com/lib/pq"
)

var (
	Connx *sql.DB
	once  sync.Once
	err   error
)

func ConnectDatabase() *sql.DB {
	once.Do(func() {

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			constans.Envs.DBHost,
			constans.Envs.DBPort,
			constans.Envs.DBUser,
			constans.Envs.DBPassword,
			constans.Envs.DBName,
		)

		Connx, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Error connecting to database", err)
		}

		err = Connx.Ping()
		if err != nil {
			log.Fatal("Error pinging database", err)
		}

		Connx.SetConnMaxIdleTime(5 * time.Minute) // 5 minutes
		Connx.SetMaxOpenConns(20)
		Connx.SetMaxIdleConns(10)

		log.Println("Postgres database connected")
	})

	return Connx
}
