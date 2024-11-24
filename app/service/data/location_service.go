package data

import (
	"factory_management_go/app/domain/dao/location"
	repo "factory_management_go/app/repository/data"
)

type LocationService struct {
	Repository *repo.LocationRepository
}

func (l *LocationService) Initialise() {
	l.Repository = &repo.LocationRepository{}
	l.Repository.Initialise()
}

func (l *LocationService) GetAllLocationsFromLocationType(locationTypeId string, companyId string) (*[]location.Location, error) {
	return l.Repository.FindAllByLocationTypeAndMetadataCompanyId(locationTypeId, companyId)
}
func (l *LocationService) GetAllLocations(companyId string) (*[]location.Location, error) {
	return l.Repository.FindAllByMetadataCompanyId(companyId)
}
