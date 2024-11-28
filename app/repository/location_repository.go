package repository

import (
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
	engine "factory_management_go/app/engine"
	logg "factory_management_go/app/log"
)

type LocationRepository struct {
	Eng *engine.RepoEngine[location.Location]
	Ltr *LocationTypeRepository
}

func (l *LocationRepository) FindAllByLocationTypeAndMetadataCompanyId(locationTypeId string, companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{LocationTypeId: locationTypeId, Metadata: metadata.Metadata{CompanyId: companyId}}
	result, err := l.Eng.ReadAll(locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
	l.getAssociatedData(companyId, result...)
	return result, err
}

func (l *LocationRepository) FindAllByMetadataCompanyId(companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{Metadata: metadata.Metadata{CompanyId: companyId}}
	result, err := l.Eng.ReadAll(locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
	l.getAssociatedData(companyId, result...)
	return result, err
}

func (l *LocationRepository) FindByIdMetadataCompanyId(locationId string, companyId string) (*location.Location, error) {
	var result = &location.Location{}
	locationCondition := &location.Location{Id: locationId, Metadata: metadata.Metadata{CompanyId: companyId}}
	err := l.Eng.ReadOne(result, locationCondition)
	l.getAssociatedData(companyId, result)
	return result, err
}

func (l *LocationRepository) getAssociatedData(companyId string, locations ...*location.Location) {
	var locationTypeIds = make([]string, 0)
	for _, l := range locations {
		locationTypeIds = append(locationTypeIds, l.LocationTypeId)
	}
	val, err := l.Ltr.FindAllInIdsAndMetadataCompanyId(companyId, locationTypeIds...)
	if err != nil {
		logg.Logger.Error(err.Error())
	} else {
		locationTypeMap := make(map[string]*location.LocationType, len(val))
		for _, lt := range val {
			locationTypeMap[lt.Id] = lt
		}
		for _, l := range locations {
			l.SetLocationType(locationTypeMap[l.LocationTypeId])
		}
	}
}
