package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zubroide/gorm-crud/common"
)

type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	db *gorm.DB
	logger common.LoggerInterface
}

func NewBaseRepository(db *gorm.DB, logger common.LoggerInterface) BaseRepositoryInterface {
	return &BaseRepository{db: db, logger: logger}
}
