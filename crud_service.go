package gorm_crud

type CrudServiceInterface interface {
	BaseServiceInterface
	GetModel() InterfaceEntity
	GetItem(id uint) (InterfaceEntity, error)
	GetList(parameters ListParametersInterface) ([]InterfaceEntity, error)
	Create(item InterfaceEntity) InterfaceEntity
	Update(item InterfaceEntity) InterfaceEntity
	Delete(id uint) error
}

type CrudService struct {
	*BaseService
	repository CrudRepositoryInterface
}

func NewCrudService(repository CrudRepositoryInterface, logger LoggerInterface) CrudServiceInterface {
	service := NewBaseService(repository, logger)
	return &CrudService{service, repository}
}

func (c CrudService) GetModel() InterfaceEntity {
	return c.repository.GetModel()
}

func (c CrudService) GetItem(id uint) (InterfaceEntity, error) {
	return c.repository.Find(id)
}

func (c CrudService) GetList(parameters ListParametersInterface) ([]InterfaceEntity, error) {
	return c.repository.List(parameters)
}

func (c CrudService) Create(item InterfaceEntity) InterfaceEntity {
	return c.repository.Create(item)
}

func (c CrudService) Update(item InterfaceEntity) InterfaceEntity {
	return c.repository.Update(item)
}

func (c CrudService) Delete(id uint) error {
	return c.repository.Delete(id)
}
