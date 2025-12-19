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
	datab, err := db.NewMySQLStorage(mysql.Config{
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

	initStorage(datab)

	server := api.NewAPIServer(":8080", datab)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
