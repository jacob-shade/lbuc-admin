package database

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql/v2"
	"github.com/jacobshade/lbuc-admin/server/model"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Store    *session.Store
	dbModels = append(make([]interface{}, 0),
		model.User{},
		model.Player{},
		model.Team{},
		model.Task{},
	)
)

func ConnectToDatabase() {
	// Loading in database enviornment variables.
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	address := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		username, password, protocol, address, port, dbName)

	// Connecting to database
	var err error
	DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("Connected to %v database\n", dbName)

	//Auto migrating schema to keep up to date.
	for _, model := range dbModels {
		DB.AutoMigrate(model)
	}
	fmt.Println("Automigrating database schema")
}

// Adapters peiced together from gofiber storage and gorm
func SetupSessionStore() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	storage := mysql.New(mysql.Config{
		Db:    sqlDB,
		Table: "sessions_store",
		Reset: false,
	})

	Store = session.New(session.Config{
		CookieSecure: true,
		Storage:      storage,
		Expiration:   7 * 24 * time.Hour, // 1 week
	})
}
