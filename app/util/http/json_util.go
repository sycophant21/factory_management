package http

import (
	"factory_management_go/app/domain/dao"
	"factory_management_go/app/domain/dao/component"
	"factory_management_go/app/domain/dao/location"
	dto "factory_management_go/app/domain/dto/response/data"
	dtoWrapper "factory_management_go/app/domain/dto/response/data/wrapper"
	responseWrapper "factory_management_go/app/domain/dto/response/wrapper"
	"strconv"
	"strings"
	"time"
)

func ConvertLocationToLocationResponseDtoForView(location location.Location) *dto.LocationResponseDto {
	return &dto.LocationResponseDto{ResponseDto: &responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: location.Metadata.CreationTimestamp, LastUpdatedAt: location.Metadata.LastUpdatedTimestamp}}, Id: location.Id, Name: location.Name, Details: location.Details, LocationType: location.LocationType.Name}
}
func ConvertLocationToLocationResponseDtoForEdit(location location.Location) *dto.LocationResponseDto {
	return &dto.LocationResponseDto{ResponseDto: &responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: location.Metadata.CreationTimestamp, LastUpdatedAt: location.Metadata.LastUpdatedTimestamp}}, Id: location.Id, Name: location.Name, Details: location.Details, LocationType: location.LocationType.Id}
}
func ConvertAllLocationsToLocationResponseDto(locations []*location.Location) *dtoWrapper.AllDataResponseDto {
	var data = make([]dtoWrapper.Data, 0)
	for _, lt := range locations {
		var ltd dtoWrapper.Data = ConvertLocationToLocationResponseDtoForView(*lt)
		data = append(data, ltd)
	}
	now := time.Now()
	return &dtoWrapper.AllDataResponseDto{ResponseDto: responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}}, Data: data}
}
func ConvertLocationTypeToLocationTypeResponseDto(locationType location.LocationType) *dto.LocationTypeResponseDto {
	return &dto.LocationTypeResponseDto{ResponseDto: &responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: locationType.Metadata.CreationTimestamp, LastUpdatedAt: locationType.Metadata.LastUpdatedTimestamp}}, Id: locationType.Id, Name: locationType.Name, Code: locationType.Code}
}
func ConvertAllLocationTypesToLocationTypeResponseDto(locationTypes []*location.LocationType) *dtoWrapper.AllDataResponseDto {
	var data = make([]dtoWrapper.Data, 0)
	for _, lt := range locationTypes {
		var ltd dtoWrapper.Data = ConvertLocationTypeToLocationTypeResponseDto(*lt)
		data = append(data, ltd)
	}
	now := time.Now()
	return &dtoWrapper.AllDataResponseDto{ResponseDto: responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}}, Data: data}
}

func ConvertComponentTypeToComponentTypeResponseDto(componentType component.ComponentType) *dto.ComponentTypeResponseDto {
	return &dto.ComponentTypeResponseDto{ResponseDto: &responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: componentType.Metadata.CreationTimestamp, LastUpdatedAt: componentType.Metadata.LastUpdatedTimestamp}}, Id: componentType.Id, Name: componentType.Name}
}
func ConvertAllComponentTypesToComponentTypeResponseDto(componentTypes []*component.ComponentType) *dtoWrapper.AllDataResponseDto {
	var data = make([]dtoWrapper.Data, 0)
	for _, ct := range componentTypes {
		var ctd dtoWrapper.Data = ConvertComponentTypeToComponentTypeResponseDto(*ct)
		data = append(data, ctd)
	}
	now := time.Now()
	return &dtoWrapper.AllDataResponseDto{ResponseDto: responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}}, Data: data}
}

func ConvertComponentToComponentResponseDtoForEdit(component component.Component) *dto.ComponentResponseDto {
	return &dto.ComponentResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: component.Metadata.CreationTimestamp, LastUpdatedAt: component.Metadata.LastUpdatedTimestamp,
			},
		},
		Id:                   component.Id,
		Name:                 component.Name,
		Description:          component.Description,
		ComponentType:        component.ComponentType.Code,
		ComponentUseCategory: component.ComponentInfo.ComponentUseCategory.Code,
		ComponentReturnType:  component.ComponentInfo.OldPartReturnType.Code,
		ComponentPosition:    component.ComponentInfo.OldPartPosition.Code,
		ExternalUnit:         component.ComponentInfo.ExternalUnit.Code,
		InternalUnit:         component.ComponentInfo.InternalUnit.Code,
		ExternalUnits:        component.ComponentInfo.ConversionFactor.ExternalUnits,
		InternalUnits:        component.ComponentInfo.ConversionFactor.InternalUnits,
		ConversionFactor:     strconv.Itoa(int(component.ComponentInfo.ConversionFactor.ExternalUnits)) + ":" + strconv.Itoa(int(component.ComponentInfo.ConversionFactor.InternalUnits)),
		HsnCode:              component.ComponentInfo.HsnCode,
		GstPercentage:        component.ComponentInfo.GstPercentageId,
	}
}

func ConvertComponentToComponentResponseDtoForView(component component.Component) *dto.ComponentResponseDto {
	return &dto.ComponentResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: component.Metadata.CreationTimestamp, LastUpdatedAt: component.Metadata.LastUpdatedTimestamp,
			},
		},
		Id:                   component.Id,
		Name:                 component.Name,
		Description:          component.Description,
		ComponentType:        component.ComponentType.Name,
		ComponentUseCategory: component.ComponentInfo.ComponentUseCategory.Name,
		ComponentReturnType:  component.ComponentInfo.OldPartReturnType.Name,
		ComponentPosition:    component.ComponentInfo.OldPartPosition.Name,
		ExternalUnit:         component.ComponentInfo.ExternalUnit.Name,
		InternalUnit:         component.ComponentInfo.InternalUnit.Name,
		ExternalUnits:        component.ComponentInfo.ConversionFactor.ExternalUnits,
		InternalUnits:        component.ComponentInfo.ConversionFactor.InternalUnits,
		ConversionFactor:     strconv.FormatUint(uint64(component.ComponentInfo.ConversionFactor.ExternalUnits), 10) + ":" + strconv.FormatUint(uint64(component.ComponentInfo.ConversionFactor.InternalUnits), 10),
		HsnCode:              component.ComponentInfo.HsnCode,
		GstPercentage:        strconv.FormatFloat(float64(component.ComponentInfo.GstPercentage.Value), 'f', 5, 32) + " %",
	}
}

func ConvertAllComponentsToComponentResponseDto(components []*component.Component) *dtoWrapper.AllDataResponseDto {
	var data = make([]dtoWrapper.Data, 0)
	for _, c := range components {
		var cd dtoWrapper.Data = ConvertComponentToComponentResponseDtoForView(*c)
		data = append(data, cd)
	}
	now := time.Now()
	return &dtoWrapper.AllDataResponseDto{ResponseDto: responseWrapper.ResponseDto{Metadata: &responseWrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}}, Data: data}
}

func ConvertComponentInventoryToComponentInventoryResponseDtoForEdit(componentInventory component.ComponentInventory) *dto.ComponentInventoryResponseDto {
	return &dto.ComponentInventoryResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: componentInventory.Metadata.CreationTimestamp, LastUpdatedAt: componentInventory.Metadata.LastUpdatedTimestamp,
			},
		},
		Id:                        componentInventory.Id,
		Name:                      componentInventory.Name,
		PrintName:                 componentInventory.PrintName,
		Description:               componentInventory.Description,
		Note:                      componentInventory.Note,
		Component:                 componentInventory.Component.Id,
		ComponentUseCategory:      componentInventory.Component.ComponentInfo.ComponentUseCategory.Code,
		ComponentReturnType:       componentInventory.Component.ComponentInfo.OldPartReturnType.Code,
		ComponentPosition:         componentInventory.Component.ComponentInfo.OldPartPosition.Code,
		ExternalUnit:              componentInventory.ComponentInventoryInfo.ExternalUnit.Code,
		InternalUnit:              componentInventory.ComponentInventoryInfo.InternalUnit.Code,
		ExternalUnits:             componentInventory.ComponentInventoryInfo.ConversionFactor.ExternalUnits,
		InternalUnits:             componentInventory.ComponentInventoryInfo.ConversionFactor.InternalUnits,
		ConversionFactor:          strconv.FormatUint(uint64(componentInventory.ComponentInventoryInfo.ConversionFactor.ExternalUnits), 10) + ":" + strconv.FormatUint(uint64(componentInventory.ComponentInventoryInfo.ConversionFactor.InternalUnits), 10),
		HsnCode:                   componentInventory.ComponentInventoryInfo.HsnCode,
		GstPercentage:             componentInventory.ComponentInventoryInfo.GstPercentage.Id,
		Location:                  componentInventory.ComponentStockInformation.Location.Id,
		MaintainBatches:           strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.MaintainBatches)),
		MinStock:                  componentInventory.ComponentStockInformation.MinStock,
		MaxStock:                  componentInventory.ComponentStockInformation.MaxStock,
		CurrentStock:              componentInventory.ComponentStockInformation.CurrentStock,
		HaveWarranty:              strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.HaveWarranty)),
		ReorderQuantity:           componentInventory.ComponentStockInformation.ReorderQuantity,
		StandardPrice:             componentInventory.ComponentStockInformation.StandardPrice,
		Active:                    strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.Active)),
		ComponentLifeDurationType: componentInventory.ComponentStockInformation.Life.DurationType.Code,
		ComponentLifeDuration:     componentInventory.ComponentStockInformation.Life.Duration,
	}
}

func ConvertComponentInventoryToComponentInventoryResponseDtoForView(componentInventory component.ComponentInventory) *dto.ComponentInventoryResponseDto {
	return &dto.ComponentInventoryResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: componentInventory.Metadata.CreationTimestamp, LastUpdatedAt: componentInventory.Metadata.LastUpdatedTimestamp,
			},
		},
		Id:                        componentInventory.Id,
		Name:                      componentInventory.Name,
		PrintName:                 componentInventory.PrintName,
		Description:               componentInventory.Description,
		Note:                      componentInventory.Note,
		Component:                 componentInventory.Component.Name,
		ComponentUseCategory:      componentInventory.Component.ComponentInfo.ComponentUseCategory.Name,
		ComponentReturnType:       componentInventory.Component.ComponentInfo.OldPartReturnType.Name,
		ComponentPosition:         componentInventory.Component.ComponentInfo.OldPartPosition.Name,
		ExternalUnit:              componentInventory.ComponentInventoryInfo.ExternalUnit.Name,
		InternalUnit:              componentInventory.ComponentInventoryInfo.InternalUnit.Name,
		ExternalUnits:             componentInventory.ComponentInventoryInfo.ConversionFactor.ExternalUnits,
		InternalUnits:             componentInventory.ComponentInventoryInfo.ConversionFactor.InternalUnits,
		ConversionFactor:          strconv.FormatUint(uint64(componentInventory.ComponentInventoryInfo.ConversionFactor.ExternalUnits), 10) + ":" + strconv.FormatUint(uint64(componentInventory.ComponentInventoryInfo.ConversionFactor.InternalUnits), 10),
		HsnCode:                   componentInventory.ComponentInventoryInfo.HsnCode,
		GstPercentage:             strconv.FormatFloat(float64(componentInventory.ComponentInventoryInfo.GstPercentage.Value), 'f', 5, 32) + " %",
		Location:                  componentInventory.ComponentStockInformation.Location.Name,
		MaintainBatches:           strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.MaintainBatches)),
		MinStock:                  componentInventory.ComponentStockInformation.MinStock,
		MaxStock:                  componentInventory.ComponentStockInformation.MaxStock,
		CurrentStock:              componentInventory.ComponentStockInformation.CurrentStock,
		HaveWarranty:              strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.HaveWarranty)),
		ReorderQuantity:           componentInventory.ComponentStockInformation.ReorderQuantity,
		StandardPrice:             componentInventory.ComponentStockInformation.StandardPrice,
		Active:                    strings.ToUpper(strconv.FormatBool(componentInventory.ComponentStockInformation.Active)),
		ComponentLifeDurationType: componentInventory.ComponentStockInformation.Life.DurationType.Name,
		ComponentLifeDuration:     componentInventory.ComponentStockInformation.Life.Duration,
	}
}

func ConvertComponentUseCategoryToComponentUseCategoryResponseDto(componentUseCategory component.ComponentUseCategory) *dto.ComponentUseCategoryResponseDto {
	return &dto.ComponentUseCategoryResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: componentUseCategory.Metadata.CreationTimestamp, LastUpdatedAt: componentUseCategory.Metadata.LastUpdatedTimestamp,
			},
		},
		Id: componentUseCategory.Id, Name: componentUseCategory.Name,
	}
}

func ConvertComponentReturnTypeToComponentReturnTypeResponseDto(componentReturnType component.ComponentReturnType) *dto.ComponentReturnTypeResponseDto {
	return &dto.ComponentReturnTypeResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: componentReturnType.Metadata.CreationTimestamp, LastUpdatedAt: componentReturnType.Metadata.LastUpdatedTimestamp,
			},
		},
		Id: componentReturnType.Id, Name: componentReturnType.Name,
	}
}

func ConvertComponentPositionToComponentPositionResponseDto(componentPosition component.ComponentPosition) *dto.ComponentPositionResponseDto {
	return &dto.ComponentPositionResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: componentPosition.Metadata.CreationTimestamp, LastUpdatedAt: componentPosition.Metadata.LastUpdatedTimestamp,
			},
		},
		Id: componentPosition.Id, Name: componentPosition.Name,
	}
}

func ConvertGstPercentageToGstPercentageResponseDto(gstPercentage dao.GSTPercentage) *dto.GSTPercentageResponseDto {
	return &dto.GSTPercentageResponseDto{
		ResponseDto: &responseWrapper.ResponseDto{
			Metadata: &responseWrapper.ResponseMetadata{
				Message: "Success", HttpCode: 200, CreatedAt: gstPercentage.Metadata.CreationTimestamp, LastUpdatedAt: gstPercentage.Metadata.LastUpdatedTimestamp,
			},
		},
		Id: gstPercentage.Id, Value: gstPercentage.Value,
	}
}
