package gorm_crud

import (
	"github.com/jinzhu/gorm"
)

type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	Db     *gorm.DB
	Logger LoggerInterface
}

func NewBaseRepository(db *gorm.DB, logger LoggerInterface) BaseRepositoryInterface {
	return &BaseRepository{Db: db, Logger: logger}
}
