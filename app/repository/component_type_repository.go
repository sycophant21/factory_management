package repository

import (
	component "factory_management_go/app/domain/dao/component"
	metadata "factory_management_go/app/domain/metadata"
	engine "factory_management_go/app/engine"
)

type ComponentTypeRepository struct {
	Eng *engine.RepoEngine[component.ComponentType]
}

func (ctr *ComponentTypeRepository) FindAllInIdsAndMetadataCompanyId(companyId string, componentTypeIds ...string) ([]*component.ComponentType, error) {
	var results []*component.ComponentType
	var conditions []*component.ComponentType = make([]*component.ComponentType, 0)
	for _, componentTypeId := range componentTypeIds {
		conditions = append(conditions, &component.ComponentType{Id: componentTypeId, Metadata: &metadata.Metadata{CompanyId: companyId}})
	}
	results, err := ctr.Eng.ReadAll(conditions...)
	return results, err
}

func (ctr *ComponentTypeRepository) FindAllByMetadataCompanyId(companyId string) ([]*component.ComponentType, error) {
	var result []*component.ComponentType
	componentTypeCondition := &component.ComponentType{Metadata: &metadata.Metadata{CompanyId: companyId}}
	result, err := ctr.Eng.ReadAll(componentTypeCondition)
	if err != nil {
		r := make([]*component.ComponentType, 0)
		result = r
	}
	return result, err
}

func (ctr *ComponentTypeRepository) FindByIdAndMetadataCompanyId(componentTypeId string, companyId string) (*component.ComponentType, error) {
	var result = &component.ComponentType{}
	componentTypeCondition := &component.ComponentType{Id: componentTypeId, Metadata: &metadata.Metadata{CompanyId: companyId}}
	err := ctr.Eng.ReadOne(result, componentTypeCondition)
	return result, err
}
