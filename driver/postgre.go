package driver

import (
	"fmt"
	"redigo-example/config"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbConfig config.DBConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode = disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	// db.Migrator().DropTable(entity.User{})
	// db.Migrator().DropTable(entity.Book{})
	// db.AutoMigrate(entity.User{})
	// db.AutoMigrate(entity.Book{})
	// entity.UserSeeder(db)
	// entity.BookSeeder(db)

	return
}
