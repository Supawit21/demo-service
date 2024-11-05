package database

import (
	"Supawit21/demo_service/pkg/utils"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialDatabase() *gorm.DB {
	dsn, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		panic("fatal error postgres url")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fatal error postgres connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("fatal error postgres connection pool")
	}

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_IDLE"))

	sqlDB.SetMaxIdleConns(maxConn)
	sqlDB.SetMaxOpenConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Second)

	return db
}
