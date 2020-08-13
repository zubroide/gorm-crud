package service

import (
	"github.com/zubroide/gorm-crud/common"
	"github.com/zubroide/gorm-crud/repository"
)

type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	repository repository.BaseRepositoryInterface
	logger common.LoggerInterface
}

func NewBaseService(repository repository.BaseRepositoryInterface, logger common.LoggerInterface) *BaseService {
	return &BaseService{repository: repository, logger: logger}
}
