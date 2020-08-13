package gorm_crud

type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	repository BaseRepositoryInterface
	logger     LoggerInterface
}

func NewBaseService(repository BaseRepositoryInterface, logger LoggerInterface) *BaseService {
	return &BaseService{repository: repository, logger: logger}
}
