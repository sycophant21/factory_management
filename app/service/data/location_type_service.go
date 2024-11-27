package data

import (
	"factory_management_go/app/domain/dao/location"
	repo "factory_management_go/app/repository/data"
)

type LocationTypeService struct {
	Repository *repo.LocationTypeRepository
}

func (l *LocationTypeService) GetAllLocationTypes(companyId string) ([]*location.LocationType, error) {
	return l.Repository.FindAllByMetadataCompanyId(companyId)
}
func (l *LocationTypeService) GetLocationTypeDetails(locationTypeId string, companyId string) (location.LocationType, error) {
	val, err := l.Repository.FindByIdMetadataCompanyId(locationTypeId, companyId)
	return *val, err
}
func (l *LocationTypeService) GetLocationTypeDetailsFromCode(locationTypeCode string, companyId string) (location.LocationType, error) {
	val, err := l.Repository.FindByCodeMetadataCompanyId(locationTypeCode, companyId)
	return *val, err
}
