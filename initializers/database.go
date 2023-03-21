package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	dsn := "host=tiny.db.elephantsql.com user=yajuioky password=3gEAqbRprCC3U1dkRrXACdK_snqT_pGR dbname=yajuioky port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}
}
