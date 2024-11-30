package option

import (
	optionDto "factory_management_go/app/domain/dto/response/option"
	logg "factory_management_go/app/log"
	service "factory_management_go/app/service/data"
)

type OptionsService struct {
	LocationService      *service.LocationService
	LocationTypeService  *service.LocationTypeService
	ComponentTypeService *service.ComponentTypeService
}

func (o *OptionsService) GetAllLocationTypeOptions(companyId string) ([]*optionDto.OptionWrapper, error) {
	var options []*optionDto.OptionWrapper = make([]*optionDto.OptionWrapper, 0)
	locationTypes, err := o.LocationTypeService.GetAllLocationTypes(companyId)
	if err != nil {
		logg.Logger.Error("Error while fetching LocationTypes with Company Id: {" + companyId + "},\nError: " + err.Error())
	} else {
		for index, lt := range locationTypes {
			var option optionDto.OptionWrapper = optionDto.OptionWrapper{Option: lt, Index: uint8(index + 1)}
			options = append(options, &option)
		}
	}
	return options, err
}
func (o *OptionsService) GetAllComponentTypeOptions(companyId string) ([]*optionDto.OptionWrapper, error) {
	var options []*optionDto.OptionWrapper = make([]*optionDto.OptionWrapper, 0)
	componentTypes, err := o.ComponentTypeService.GetAllComponentTypes(companyId)
	if err != nil {
		logg.Logger.Error("Error while fetching LocationTypes with Company Id: {" + companyId + "},\nError: " + err.Error())
	} else {
		for index, c := range componentTypes {
			var option optionDto.OptionWrapper = optionDto.OptionWrapper{Option: c, Index: uint8(index + 1)}
			options = append(options, &option)
		}
	}
	return options, err
}

/*
List<LocationType> locationTypes = locationTypeService.getAllLocationTypes(companyId);
List<EntityOptionResponseDto> locationTypeResponseDtos = new ArrayList<>();
int index = 1;
for (LocationType locationType : locationTypes) {
    locationTypeResponseDtos.add(DaoToDto.convertLocationTypeToEntityOptionResponseDto(index++, locationType));
}
return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), locationTypeResponseDtos);
*/
