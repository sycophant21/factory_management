package repository

import (
	component "factory_management_go/app/domain/dao/component"
	metadata "factory_management_go/app/domain/metadata"
	engine "factory_management_go/app/engine"
	logg "factory_management_go/app/log"
)

type ComponentRepository struct {
	Eng                     *engine.RepoEngine[component.Component]
	ComponentTypeRepository *ComponentTypeRepository
}

func (cr *ComponentRepository) FindAllByMetadataCompanyId(companyId string) ([]*component.Component, error) {
	var result []*component.Component
	componentCondition := &component.Component{Metadata: &metadata.Metadata{CompanyId: companyId}}
	result, err := cr.Eng.ReadAll(componentCondition)
	if err != nil {
		r := make([]*component.Component, 0)
		result = r
	}
	cr.getAssociatedData(companyId, result...)
	return result, err
}

func (cr *ComponentRepository) FindByIdMetadataCompanyId(componentId string, companyId string) (*component.Component, error) {
	var result = &component.Component{}
	locationCondition := &component.Component{Id: componentId, Metadata: &metadata.Metadata{CompanyId: companyId}}
	err := cr.Eng.ReadOne(result, locationCondition)
	cr.getAssociatedData(companyId, result)
	return result, err
}

func (cr *ComponentRepository) getAssociatedData(companyId string, components ...*component.Component) {
	var componentTypeIds = make([]string, 0)
	for _, c := range components {
		componentTypeIds = append(componentTypeIds, c.ComponentTypeId)
	}
	val, err := cr.ComponentTypeRepository.FindAllInIdsAndMetadataCompanyId(companyId, componentTypeIds...)
	if err != nil {
		logg.Logger.Error(err.Error())
	} else {
		componentTypeMap := make(map[string]*component.ComponentType, len(val))
		for _, ct := range val {
			componentTypeMap[ct.Id] = ct
		}
		for _, c := range components {
			c.ComponentType = componentTypeMap[c.ComponentTypeId]
			//l.SetLocationType(componentTypeMap[l.LocationTypeId])
		}
	}
}
