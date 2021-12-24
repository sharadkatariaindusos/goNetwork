package http

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	PUT    HttpMethod = "PUT"
	POST   HttpMethod = "POST"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)
