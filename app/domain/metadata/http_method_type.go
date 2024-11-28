package metadata

type HttpMethodType string

const (
	GET     HttpMethodType = "GET"
	HEAD    HttpMethodType = "HEAD"
	POST    HttpMethodType = "POST"
	PUT     HttpMethodType = "PUT"
	PATCH   HttpMethodType = "PATCH"
	DELETE  HttpMethodType = "DELETE"
	OPTIONS HttpMethodType = "OPTIONS"
	TRACE   HttpMethodType = "TRACE"
)
