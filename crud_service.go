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
	Repository CrudRepositoryInterface
}

func NewCrudService(repository CrudRepositoryInterface, logger LoggerInterface) CrudServiceInterface {
	service := NewBaseService(repository, logger)
	return &CrudService{service, repository}
}

func (c CrudService) GetModel() InterfaceEntity {
	return c.Repository.GetModel()
}

func (c CrudService) GetItem(id uint) (InterfaceEntity, error) {
	return c.Repository.Find(id)
}

func (c CrudService) GetList(parameters ListParametersInterface) ([]InterfaceEntity, error) {
	return c.Repository.List(parameters)
}

func (c CrudService) Create(item InterfaceEntity) InterfaceEntity {
	return c.Repository.Create(item)
}

func (c CrudService) Update(item InterfaceEntity) InterfaceEntity {
	return c.Repository.Update(item)
}

func (c CrudService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
