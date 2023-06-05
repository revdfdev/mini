package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func (r *Response) StatusCode(code int) *Response {
	r.WriteHeader(code)
	return r
}

func (r *Response) SetHeader(key, value string) *Response {
	r.Header().Set(key, value)
	return r
}

func (r *Response) JSON(data interface{}) (*Response, error) {
	if err := json.NewEncoder(r.ResponseWriter).Encode(data); err != nil {
		return nil, err
	}

	return r, nil
}

// func (response *Response) StatusCode(statusCode int) {
// 	response.WriteHeader(statusCode)
// }

// func (response *Response) SetHeader(key, value string) {
// 	response.Header().Set(key, value)
// }

// func (response *Response) JSON(data interface{}) error {
// 	response.SetHeader("Content-Type", "application/json")
// 	if err := json.NewEncoder(response.ResponseWriter).Encode(data); err != nil {
// 		return err
// 	}

// 	return nil
// }
