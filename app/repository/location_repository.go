package repository

import (
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
	engine "factory_management_go/app/engine"
	logg "factory_management_go/app/log"
)

type LocationRepository struct {
	Eng                    *engine.RepoEngine[location.Location]
	LocationTypeRepository *LocationTypeRepository
}

func (lr *LocationRepository) FindAllByLocationTypeIdAndMetadataCompanyId(locationTypeId string, companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{LocationTypeId: locationTypeId, Metadata: &metadata.Metadata{CompanyId: companyId}}
	result, err := lr.Eng.ReadAll(locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
	lr.getAssociatedData(companyId, result...)
	return result, err
}

func (lr *LocationRepository) FindAllByLocationTypeCodeAndMetadataCompanyId(locationTypeCode string, companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{Metadata: &metadata.Metadata{CompanyId: companyId}}

	result, err := lr.Eng.ReadAllFromNestedJoinField(engine.JoinClause{JoinType: engine.INNER, TableName: (&location.LocationType{}).TableName(), Condition: "location.location_type_id = location_type.id", Args: nil}, []engine.WhereClause{{ParamName: "code", ParamValue: locationTypeCode}}, &location.Location{}, locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
	lr.getAssociatedData(companyId, result...)
	return result, err
}

func (lr *LocationRepository) FindAllByMetadataCompanyId(companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{Metadata: &metadata.Metadata{CompanyId: companyId}}
	result, err := lr.Eng.ReadAll(locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
	lr.getAssociatedData(companyId, result...)
	return result, err
}

func (lr *LocationRepository) FindByIdAndMetadataCompanyId(locationId string, companyId string) (*location.Location, error) {
	var result = &location.Location{}
	locationCondition := &location.Location{Id: locationId, Metadata: &metadata.Metadata{CompanyId: companyId}}
	err := lr.Eng.ReadOne(result, locationCondition)
	lr.getAssociatedData(companyId, result)
	return result, err
}

func (lr *LocationRepository) getAssociatedData(companyId string, locations ...*location.Location) {
	var locationTypeIds = make([]string, 0)
	for _, l := range locations {
		locationTypeIds = append(locationTypeIds, l.LocationTypeId)
	}
	val, err := lr.LocationTypeRepository.FindAllInIdsAndMetadataCompanyId(companyId, locationTypeIds...)
	if err != nil {
		logg.Logger.Error(err.Error())
	} else {
		locationTypeMap := make(map[string]*location.LocationType, len(val))
		for _, lt := range val {
			locationTypeMap[lt.Id] = lt
		}
		for _, l := range locations {
			l.LocationType = locationTypeMap[l.LocationTypeId]
			//l.SetLocationType(locationTypeMap[l.LocationTypeId])
		}
	}
}
