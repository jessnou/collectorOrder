package db

import (
	"collectorOrder/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	dbConn *sqlx.DB
)

func GetDBConn() (*sqlx.DB, error) {

	conf, err := config.LoadConfig(".")
	if err != nil {
		if err != nil {
			log.Fatal("cannot load config", err)
		}
	}

	dbConn, err = sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DbUser, conf.DbPassword, conf.DbName))
	if err != nil {
		log.Fatal("cannot open db", err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal("cannot ping db", err)
	}

	return dbConn, nil
}
