package data

import (
	component "factory_management_go/app/domain/dao/component"
	repository "factory_management_go/app/repository"
)

type ComponentService struct {
	ComponentRepository *repository.ComponentRepository
}

func (cs *ComponentService) GetAllComponents(companyId string) ([]*component.Component, error) {
	return cs.ComponentRepository.FindAllByMetadataCompanyId(companyId)
}

func (cs *ComponentService) GetComponentDetails(componentId string, companyId string) (*component.Component, error) {
	return cs.ComponentRepository.FindByIdMetadataCompanyId(componentId, companyId)
}
