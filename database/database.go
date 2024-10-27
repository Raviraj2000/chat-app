package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Raviraj2000/chat-app/database/models"
)

type DBConn struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	timezone string
}

var DB *gorm.DB

func Init() {

	d := &DBConn{
		host:     "localhost",
		port:     5432,
		user:     "root",
		password: "secret",
		dbname:   "chatapp",
		timezone: "America/Los_Angeles",
	}

	var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=%s",
		d.host, d.port, d.user, d.password, d.dbname, d.timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(&models.User{})

	return
}
