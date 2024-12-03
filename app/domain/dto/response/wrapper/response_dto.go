package wrapper

import "time"

type ResponseDto struct {
	Metadata *ResponseMetadata `json:"'metadata',omitempty"`
}

type ResponseMetadata struct {
	Message       string     `json:"'message'"`
	HttpCode      uint16     `json:"'httpCode'"`
	CreatedAt     *time.Time `json:"'createdAt'"`
	LastUpdatedAt *time.Time `json:"'updatedAt'"`
}
