package data

import (
	"factory_management_go/app/domain/dao/location"
	repo "factory_management_go/app/repository"
)

type LocationService struct {
	Repository          *repo.LocationRepository
	LocationTypeService *LocationTypeService
}

func (l *LocationService) GetAllLocationsFromLocationType(locationTypeId string, companyId string) ([]*location.Location, error) {
	return l.Repository.FindAllByLocationTypeAndMetadataCompanyId(locationTypeId, companyId)
}
func (l *LocationService) GetAllLocations(companyId string) ([]*location.Location, error) {
	return l.Repository.FindAllByMetadataCompanyId(companyId)
}
func (l *LocationService) GetLocationDetails(locationId string, companyId string) (location.Location, error) {
	val, err := l.Repository.FindByIdMetadataCompanyId(locationId, companyId)
	return *val, err
}
func (l *LocationService) GetAllLocationsFromLocationTypeCode(locationTypeCode string, companyId string) ([]*location.Location, error) {
	locationType, err := l.LocationTypeService.GetLocationTypeDetailsFromCode(locationTypeCode, companyId)
	if err != nil {
		r := make([]*location.Location, 0)
		result := r
		return result, err
	}
	return l.Repository.FindAllByLocationTypeAndMetadataCompanyId(locationType.Id, companyId)
}
