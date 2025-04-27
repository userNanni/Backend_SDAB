package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataBase struct {
	*sql.DB
}

type settingsDatabase struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func dbConn() *DataBase {
	dbSettings := setEnv()

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbSettings.host, dbSettings.port, dbSettings.user, dbSettings.password, dbSettings.dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return &DataBase{DB: db}
}

func CheckError(err error) {

	if err != nil {
		panic(err)
	}
}

func setEnv() settingsDatabase {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbSettings := settingsDatabase{
		host:     os.Getenv("host"),
		port:     os.Getenv("port"),
		user:     os.Getenv("user"),
		password: os.Getenv("password"),
		dbname:   os.Getenv("dbname"),
	}

	fmt.Println(dbSettings)

	return dbSettings
}

func (db *DataBase) CloseConn() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection closed")
}
