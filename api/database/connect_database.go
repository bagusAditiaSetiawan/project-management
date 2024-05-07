package database

import (
	"fmt"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewConnectDatabase() *gorm.DB {
	dbHost := config.Config("DB_HOST")
	dbUser := config.Config("DB_USER")
	dbName := config.Config("DB_NAME")
	dbPassword := config.Config("DB_PASSWORD")
	dbPort := config.Config("DB_PORT")
	dbSslMode := config.Config("DB_SSL_MODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSslMode,
	)
	configPostgres := postgres.Config{
		DSN: dsn,
	}

	db, err := gorm.Open(postgres.New(configPostgres), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	helpers.IfPanicHelper(err)

	sql, err := db.DB()
	helpers.IfPanicHelper(err)
	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(10)

	sql.SetConnMaxIdleTime(10 * time.Minute)
	sql.SetConnMaxLifetime(60 * time.Minute)

	// Migrate the schema
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Project{})
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.TaskPeople{})
	return db
}
