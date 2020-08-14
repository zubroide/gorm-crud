package gorm_crud

type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	Repository BaseRepositoryInterface
	Logger     LoggerInterface
}

func NewBaseService(repository BaseRepositoryInterface, logger LoggerInterface) *BaseService {
	return &BaseService{Repository: repository, Logger: logger}
}
