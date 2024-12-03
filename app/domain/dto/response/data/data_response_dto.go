package data

import "factory_management_go/app/domain/dto/response/wrapper"

type (
	ComponentResponseDto struct {
		*wrapper.ResponseDto
		Id                   string `json:"'id'"`
		Name                 string `json:"'name'"`
		Description          string `json:"'description'"`
		ComponentType        string `json:"'componentType'"`
		ComponentUseCategory string `json:"'componentUseCategory'"`
		ComponentReturnType  string `json:"'oldPartReturnType'"`
		ComponentPosition    string `json:"'oldPartPosition'"`
		ExternalUnit         string `json:"'externalUnit'"`
		InternalUnit         string `json:"'internalUnit'"`
		ExternalUnits        uint16 `json:"'externalUnits'"`
		InternalUnits        uint16 `json:"'internalUnits'"`
		ConversionFactor     string `json:"'conversionFactor'"`
		HsnCode              string `json:"'hsnCode'"`
		GstPercentage        string `json:"'gstPercentage'"`
	}

	ComponentInventoryResponseDto struct {
		*wrapper.ResponseDto
		Id                        string `json:"'id'"`
		Name                      string `json:"'name'"`
		PrintName                 string `json:"'printName'"`
		Description               string `json:"description'"`
		Note                      string `json:"'note'"`
		Component                 string `json:"'component_id'"`
		ComponentUseCategory      string `json:"'componentUseCategory'"`
		ComponentReturnType       string `json:"'oldPartReturnType'"`
		ComponentPosition         string `json:"'oldPartPosition'"`
		ExternalUnit              string `json:"'externalUnit'"`
		InternalUnit              string `json:"'internalUnit'"`
		ExternalUnits             uint16 `json:"'externalUnitQuantity'"`
		InternalUnits             uint16 `json:"'internalUnitQuantity'"`
		ConversionFactor          string `json:"'cf'"`
		HsnCode                   string `json:"'hsnCode'"`
		GstPercentage             string `json:"'gstPercentage'"`
		Location                  string `json:"'location'"`
		MaintainBatches           string `json:"'maintainBatches'"`
		MinStock                  uint32 `json:"'minStock'"`
		MaxStock                  uint32 `json:"'maxStock'"`
		CurrentStock              uint32 `json:"'currentStock'"`
		HaveWarranty              string `json:"'haveWarranty'"`
		ReorderQuantity           uint32 `json:"'reorderQuantity'"`
		StandardPrice             uint32 `json:"'standardPrice'"`
		Active                    string `json:"'active'"`
		ComponentLifeDurationType string `json:"'componentLifeDurationType'"`
		ComponentLifeDuration     uint8  `json:"'componentLifeDuration'"`
	}

	ComponentPositionResponseDto struct {
		*wrapper.ResponseDto
		Id   string `json:"'id'"`
		Name string `json:"'name'"`
	}

	ComponentReturnTypeResponseDto struct {
		*wrapper.ResponseDto
		Id   string `json:"'id'"`
		Name string `json:"'name'"`
	}

	ComponentTypeResponseDto struct {
		*wrapper.ResponseDto
		Id   string `json:"'id'"`
		Name string `json:"'name'"`
	}

	ComponentUseCategoryResponseDto struct {
		*wrapper.ResponseDto
		Id   string `json:"'id'"`
		Name string `json:"'name'"`
	}

	GSTPercentageResponseDto struct {
		*wrapper.ResponseDto
		Id    string  `json:"'id'"`
		Value float32 `json:"'value'"`
	}

	LocationTypeResponseDto struct {
		*wrapper.ResponseDto
		Id   string `json:"'id'"`
		Name string `json:"'name'"`
		Code string `json:"'code'"`
	}

	LocationResponseDto struct {
		*wrapper.ResponseDto
		Id           string `json:"'id'"`
		Name         string `json:"'name'"`
		Details      string `json:"'details'"`
		LocationType string `json:"'locationType'"`
	}
)

func (*ComponentResponseDto) IsResponseEntity() {}
func (crd *ComponentResponseDto) OmitMetadata() {
	crd.Metadata = nil
}

func (*ComponentInventoryResponseDto) IsResponseEntity() {}
func (crd *ComponentInventoryResponseDto) OmitMetadata() {
	crd.Metadata = nil
}

func (*ComponentPositionResponseDto) IsResponseEntity() {}
func (cprd *ComponentPositionResponseDto) OmitMetadata() {
	cprd.Metadata = nil
}

func (*ComponentReturnTypeResponseDto) IsResponseEntity() {}
func (crtrd *ComponentReturnTypeResponseDto) OmitMetadata() {
	crtrd.Metadata = nil
}

func (*ComponentTypeResponseDto) IsResponseEntity() {}
func (ctrd *ComponentTypeResponseDto) OmitMetadata() {
	ctrd.Metadata = nil
}

func (*ComponentUseCategoryResponseDto) IsResponseEntity() {}
func (cucrd *ComponentUseCategoryResponseDto) OmitMetadata() {
	cucrd.Metadata = nil
}

func (*LocationTypeResponseDto) IsResponseEntity() {}
func (ltrd *LocationTypeResponseDto) OmitMetadata() {
	ltrd.Metadata = nil
}

func (*LocationResponseDto) IsResponseEntity() {}
func (lrd *LocationResponseDto) OmitMetadata() {
	lrd.Metadata = nil
}
