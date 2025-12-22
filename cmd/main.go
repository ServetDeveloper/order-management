package main

import (
	"database/sql"
	"log"

	"github.com/ServetDeveloper/order-management/cmd/api"
	"github.com/ServetDeveloper/order-management/config"
	"github.com/ServetDeveloper/order-management/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	server := api.NewAPIServer(":8080", database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
