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

func (cts *ComponentTypeService) GetComponentTypeDetails(componentTypeId string, companyId string) (component.ComponentType, error) {
	val, err := cts.ComponentTypeRepository.FindByIdAndMetadataCompanyId(componentTypeId, companyId)
	return *val, err
}
