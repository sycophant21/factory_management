package repository

import (
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
	engine "factory_management_go/app/engine"
)

type LocationTypeRepository struct {
	Eng *engine.RepoEngine[location.LocationType]
}

func (ltr *LocationTypeRepository) FindAllByMetadataCompanyId(companyId string) ([]*location.LocationType, error) {
	var result []*location.LocationType
	locationTypeCondition := &location.LocationType{Metadata: &metadata.Metadata{CompanyId: companyId}}
	result, err := ltr.Eng.ReadAll(locationTypeCondition)
	if err != nil {
		r := make([]*location.LocationType, 0)
		result = r
	}
	return result, err
}

func (ltr *LocationTypeRepository) FindByIdAndMetadataCompanyId(locationTypeId string, companyId string) (*location.LocationType, error) {
	var result = &location.LocationType{}
	locationTypeCondition := &location.LocationType{Id: locationTypeId, Metadata: &metadata.Metadata{CompanyId: companyId}}
	err := ltr.Eng.ReadOne(result, locationTypeCondition)
	return result, err
}

func (ltr *LocationTypeRepository) FindByCodeAndMetadataCompanyId(locationTypeCode string, companyId string) (*location.LocationType, error) {
	var result = &location.LocationType{}
	locationTypeCondition := &location.LocationType{Code: locationTypeCode, Metadata: &metadata.Metadata{CompanyId: companyId}}
	err := ltr.Eng.ReadOne(result, locationTypeCondition)
	return result, err
}

func (ltr *LocationTypeRepository) FindAllInIdsAndMetadataCompanyId(companyId string, locationTypeIds ...string) ([]*location.LocationType, error) {
	var results []*location.LocationType
	var conditions []*location.LocationType = make([]*location.LocationType, 0)
	for _, locationTypeId := range locationTypeIds {
		conditions = append(conditions, &location.LocationType{Id: locationTypeId, Metadata: &metadata.Metadata{CompanyId: companyId}})
	}
	results, err := ltr.Eng.ReadAll(conditions...)
	return results, err
}
