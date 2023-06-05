package mini

import "net/http"

type Request struct {
	*http.Request
}

func (request *Request) GetHeader(key string) string {
	return request.Header.Get(key)
}
