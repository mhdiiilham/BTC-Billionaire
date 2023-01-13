package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToSQL(cfg *Configuration) *sql.DB {
	log.Printf("connecting to pgsql database=%s", cfg.DBHost)
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatalf("conecting to db %s failed: %v", cfg.DBHost, err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ping to db %s failed: %v", cfg.DBHost, err)
		return nil
	}

	log.Printf("connected to db %s (%s:%s) ...", cfg.DBName, cfg.DBHost, cfg.DBPort)
	return db
}
