package cmd

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initGorm(cfg Config) (*gorm.DB, error) {
	// https://github.com/github.com/pramudya3/backend/payment/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBName,
		),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(10)
	sql.SetConnMaxIdleTime(30 * time.Minute)

	return db, nil
}
