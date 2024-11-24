package data

import (
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
	"factory_management_go/app/repository"
	"xorm.io/xorm"
)

type LocationRepository struct {
	Eng *xorm.Engine
}

func (l *LocationRepository) Initialise() {
	l.Eng = repository.Eng
}

func (l *LocationRepository) FindAllByLocationTypeAndMetadataCompanyId(locationTypeId string, companyId string) (*[]location.Location, error) {
	var locations []location.Location
	locationCondition := location.Location{LocationTypeId: locationTypeId, Metadata: metadata.Metadata{CompanyId: companyId}}
	err := l.Eng.Find(&locations, &locationCondition)
	if err != nil {
		locations = make([]location.Location, 0)
		return &locations, err
	}
	return &locations, nil
}
func (l *LocationRepository) FindAllByMetadataCompanyId(companyId string) (*[]location.Location, error) {
	var locations []location.Location
	locationCondition := location.Location{Metadata: metadata.Metadata{CompanyId: companyId}}
	err := l.Eng.Find(&locations, &locationCondition)
	if err != nil {
		locations = make([]location.Location, 0)
		return &locations, err
	}
	return &locations, nil
}

/*
List<Location> findAllByLocationTypeAndMetadata_Company_Id(LocationType locationType, String companyId);

List<Location> findAllByLocationType_CodeAndMetadata_Company_Id(String locationType, String companyId);

List<Location> findAllByMetadata_Company_Id(String companyId);

Optional<Location> findByIdAndMetadata_Company_Id(String id, String companyId);
*/
