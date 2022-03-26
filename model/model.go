package model

import (
	"GitHunter/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init the model mapping
func Init() {
	DBSection, err := config.Config.GetSection("DB")
	if err != nil {
		panic(err)
	}

	host := config.ParseKey(DBSection, "HOST")
	user := config.ParseKey(DBSection, "USER")
	password := config.ParseKey(DBSection, "PASSWORD")
	dbname := config.ParseKey(DBSection, "DBNAME")
	port := config.ParseKey(DBSection, "PORT")

	// connect to postgres DB according to DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	// migrate schema
	err = DB.AutoMigrate(&Repo{})
	if err != nil {
		panic(err)
	}
}
