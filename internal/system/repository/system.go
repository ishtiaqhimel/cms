package repository

import (
	"time"

	"gorm.io/gorm"
)

type systemRepository struct {
	*gorm.DB
}

// NewSystemRepository will create an object that represent the SystemRepository interface
func NewSystemRepository(db *gorm.DB) SystemRepository {
	return &systemRepository{
		DB: db,
	}
}

type SystemRepository interface {
	DBCheck() (bool, error)
	CurrentTime() int64
}

func (db *systemRepository) DBCheck() (bool, error) {
	sqlDB, err := db.DB.DB()
	if err == nil {
		if err = sqlDB.Ping(); err == nil {
			return true, nil
		}
	}

	return false, err
}

func (db *systemRepository) CurrentTime() int64 {
	return time.Now().Unix()
}
