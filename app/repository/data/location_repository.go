package data

import (
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
	"factory_management_go/app/repository"
)

type LocationRepository struct {
	Eng *repository.RepoEngine[location.Location]
}

func (l *LocationRepository) FindAllByLocationTypeAndMetadataCompanyId(locationTypeId string, companyId string) ([]*location.Location, error) {
	var result []*location.Location
	locationCondition := &location.Location{LocationTypeId: locationTypeId, Metadata: metadata.Metadata{CompanyId: companyId}}
	result, err := l.Eng.ReadAll(locationCondition)
	if err != nil {
		r := make([]*location.Location, 0)
		result = r
	}
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
	return result, err
}

func (l *LocationRepository) FindByIdMetadataCompanyId(locationId string, companyId string) (*location.Location, error) {
	var result = &location.Location{}
	locationCondition := &location.Location{Id: locationId, Metadata: metadata.Metadata{CompanyId: companyId}}
	err := l.Eng.ReadOne(result, locationCondition)
	return result, err
}
