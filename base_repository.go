package gorm_crud

import (
	"github.com/jinzhu/gorm"
)

type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	db     *gorm.DB
	logger LoggerInterface
}

func NewBaseRepository(db *gorm.DB, logger LoggerInterface) BaseRepositoryInterface {
	return &BaseRepository{db: db, logger: logger}
}
