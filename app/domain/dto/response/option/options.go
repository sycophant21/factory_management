package option

import (
	"encoding/json"
	"factory_management_go/app/domain/dto/response/wrapper"
)

type (
	Option interface {
		GetLabel() string
		GetValue() string
	}
	OptionWrapper struct {
		Option
		*wrapper.ResponseMetadata
		Index uint8
	}
	AllOptionsResponseDto struct {
		*wrapper.ResponseMetadata
		Data []*OptionWrapper `json:"'data'"`
	}
)

func (aos *AllOptionsResponseDto) MarshalJSON() ([]byte, error) {
	for i := range aos.Data {
		aos.Data[i].ResponseMetadata = nil
	}
	return json.Marshal(*aos)
}
func (ow OptionWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Index    uint8                     `json:"index"`
		Label    string                    `json:"label"`
		Value    string                    `json:"value"`
		Metadata *wrapper.ResponseMetadata `json:"'metadata',omitempty"`
	}{
		Index:    ow.Index,
		Label:    ow.GetLabel(),
		Value:    ow.GetValue(),
		Metadata: ow.ResponseMetadata,
	})
}
