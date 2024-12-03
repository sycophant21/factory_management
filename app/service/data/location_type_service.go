package data

import (
	"factory_management_go/app/domain/dao/location"
	repository "factory_management_go/app/repository"
)

type LocationTypeService struct {
	LocationTypeRepository *repository.LocationTypeRepository
}

func (lts *LocationTypeService) GetAllLocationTypes(companyId string) ([]*location.LocationType, error) {
	return lts.LocationTypeRepository.FindAllByMetadataCompanyId(companyId)
}
func (lts *LocationTypeService) GetLocationTypeDetails(locationTypeId string, companyId string) (location.LocationType, error) {
	val, err := lts.LocationTypeRepository.FindByIdAndMetadataCompanyId(locationTypeId, companyId)
	return *val, err
}
func (lts *LocationTypeService) GetLocationTypeDetailsFromCode(locationTypeCode string, companyId string) (location.LocationType, error) {
	val, err := lts.LocationTypeRepository.FindByCodeAndMetadataCompanyId(locationTypeCode, companyId)
	return *val, err
}
