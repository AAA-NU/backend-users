package storage

import (
	"fmt"

	"github.com/aaanu/backendusers/internal/config"
	"github.com/aaanu/backendusers/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	cfg := config.Config().Storage

	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.Name,
			cfg.Port,
			cfg.SslMode,
		),
	),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	return &Storage{db: db}
}
