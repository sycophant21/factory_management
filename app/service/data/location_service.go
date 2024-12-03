package data

import (
	"factory_management_go/app/domain/dao/location"
	repository "factory_management_go/app/repository"
)

type LocationService struct {
	LocationRepository *repository.LocationRepository
}

func (ls *LocationService) GetAllLocationsFromLocationTypeId(locationTypeId string, companyId string) ([]*location.Location, error) {
	return ls.LocationRepository.FindAllByLocationTypeIdAndMetadataCompanyId(locationTypeId, companyId)
}
func (ls *LocationService) GetAllLocations(companyId string) ([]*location.Location, error) {
	return ls.LocationRepository.FindAllByMetadataCompanyId(companyId)
}
func (ls *LocationService) GetLocationDetails(locationId string, companyId string) (location.Location, error) {
	val, err := ls.LocationRepository.FindByIdAndMetadataCompanyId(locationId, companyId)
	return *val, err
}
func (ls *LocationService) GetAllLocationsFromLocationTypeCode(locationTypeCode string, companyId string) ([]*location.Location, error) {
	return ls.LocationRepository.FindAllByLocationTypeCodeAndMetadataCompanyId(locationTypeCode, companyId)
}
