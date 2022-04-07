package config

import (
	"backend/models/global"
	"backend/models/hrd"
	"backend/models/inventory"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to create a connection to DB")
	} else {
		log.Println("Connection Established to DB")
	}

	db.AutoMigrate(&global.COMPANY{},
		&global.DEPARTEMENT{},
		&global.DIVISION{},
		&global.GLOBAL{},
		&global.LOG{},
		&global.MODULE{},
		&global.PARTNER{},
		&global.PERMISSION{},
		&global.SYSTEMSETTING{},
		&global.USER{},
		&global.USERGROUP{},
		// &inventory.CASE{},
		&inventory.CONTAINER{},
		&inventory.ITEM{},
		&inventory.LOCATION{},
		// &inventory.OBORDTL{},
		// &inventory.OBORHDR{},
		// &inventory.TASKDTL{},
		// &inventory.TASKHDR{},
		&hrd.EMPLOYEE{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from DB")
	}
	dbSQL.Close()
}
