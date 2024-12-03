package wrapper

import (
	"encoding/json"
	response "factory_management_go/app/domain/dto/response/wrapper"
)

type AllDataResponseDto struct {
	response.ResponseDto
	Data []Data
}

func (adrd *AllDataResponseDto) MarshalJSON() ([]byte, error) {
	for _, d := range adrd.Data {
		d.OmitMetadata()
	}
	return json.Marshal(*adrd)
}

type Data interface {
	IsResponseEntity()
	OmitMetadata()
}
