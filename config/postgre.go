package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializerPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.ErrF("sql opening error: %v", err)
		return nil, err
	}

	// Return Database
	return db, nil
}
