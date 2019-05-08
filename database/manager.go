package database

import (
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/xlab/closer"
	"os"
	"time"
)

const (
	maxConnections = 3
	acquireTimeout = 3 * time.Second
	pathConfig     = "/Users/mac/Desktop/TechnoparkDB-FORUM_GO/config.json"
)

type dbConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func loadConfiguration(file string) (pgxConfig pgx.ConnConfig) {
	configFile, err := os.Open(file)
	if err != nil {
		logger.Fatal.Println(err.Error())
		return
	}
	var config dbConfig
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		logger.Fatal.Println(err.Error())
		return
	}
	err = configFile.Close()
	if err != nil {
		logger.Fatal.Println(err.Error())
		return
	}

	pgxConfig.Host = config.Host
	pgxConfig.User = config.User
	pgxConfig.Password = config.Password
	pgxConfig.Database = config.DBName
	pgxConfig.Port = config.Port
	return
}

//func loadConfiguration(file string) (config dbConfig) {
//	configFile, err := os.Open(file)
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//		return
//	}
//	jsonParser := json.NewDecoder(configFile)
//	err = jsonParser.Decode(&config)
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//		return
//	}
//	err = configFile.Close()
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//		return
//	}
//
//	//pgxConfig.Host = config.Host
//	//pgxConfig.User = config.User
//	//pgxConfig.Password = config.Password
//	//pgxConfig.Database = config.DBName
//	//pgxConfig.Port = config.Port
//	return
//}

type databaseManager struct {
	dataBase *pgx.ConnPool
}

//type databaseManager struct {
//	dataBase *sql.DB
//}

var database *databaseManager

func init() {
	pgxConfig := loadConfiguration(pathConfig)
	pgxConnPoolConfig := pgx.ConnPoolConfig{ConnConfig: pgxConfig, MaxConnections: maxConnections, AcquireTimeout: acquireTimeout}

	dataBase, err := pgx.NewConnPool(pgxConnPoolConfig)
	if err != nil {
		logger.Fatal.Println(err.Error())
		return
	}

	fmt.Println("DB connection opened")

	database = &databaseManager{
		dataBase: dataBase,
	}

	closer.Bind(closeConnection)
}

func closeConnection() {
	database.dataBase.Close()
	fmt.Println("DB connection closed")
}

//func init() {
//
//	config := loadConfiguration(pathConfig)
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		config.Host, config.Port, config.User, config.Password, config.DBName)
//
//	dataBase, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//		panic(err)
//	}
//
//	err = dataBase.Ping()
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//		panic(err)
//	}
//
//	database = &databaseManager{
//		dataBase: dataBase,
//	}
//	logger.Info.Printf("\nSuccessfully connected to database at: 5432")
//
//	closer.Bind(closeConnection)
//}

//func closeConnection() {
//	err := database.dataBase.Close()
//	if err != nil {
//		logger.Fatal.Println(err.Error())
//	}
//}

func GetInstance() *databaseManager {
	return database
}
