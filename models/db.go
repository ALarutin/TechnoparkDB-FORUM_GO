package models

import (
	"data_base/presentation/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/xlab/closer"
)

const (
	Host       = "localhost"
	Port       = 5432
	Subscriber = "postgres"
	Password   = "1209qawsed"
	DBName     = "postgres"
)

type dbManager struct {
	dataBase *sql.DB
}

var db *dbManager

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, Subscriber, Password, DBName)

	dataBase, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatal.Println(err.Error())
		panic(err)
	}

	err = dataBase.Ping()
	if err != nil {
		logger.Fatal.Println(err.Error())
		panic(err)
	}

	db = &dbManager{
		dataBase: dataBase,
	}
	logger.Info.Printf("\nSuccessfully connected to database at: 5432")

	closer.Bind(closeConnection)
}

func closeConnection() {
	err := db.dataBase.Close()
	if err != nil {
		logger.Fatal.Println(err.Error())
	}
}

func GetInstance() *dbManager {
	return db
}
