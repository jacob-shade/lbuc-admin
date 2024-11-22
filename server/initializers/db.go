package initializers

import (
	"fmt"
	"lbuc-admin/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	dbModels = append(make([]interface{}, 0), models.User{}, models.Player{}, models.Team{})
)

// Connect connects to the database with the config given in the .env file and
// the models given in the database package.
//
// For details on dsn(data source name) formating, refer to [go-sql-driver]
// docs.
//
// [go-sql-driver]: https://github.com/go-sql-driver/mysql#dsn-data-source-name
func ConnectToDatabase() {
	// Loading in database enviornment variables.
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		username, password, protocol, address, port, dbName)

	// Connecting to database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
