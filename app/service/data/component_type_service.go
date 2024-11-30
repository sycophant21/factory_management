package data

import (
	"factory_management_go/app/domain/dao/component"
	"factory_management_go/app/repository"
)

type ComponentTypeService struct {
	ComponentTypeRepository *repository.ComponentTypeRepository
}

func (cts *ComponentTypeService) GetAllComponentTypes(companyId string) ([]*component.ComponentType, error) {
	return cts.ComponentTypeRepository.FindAllByMetadataCompanyId(companyId)
}
