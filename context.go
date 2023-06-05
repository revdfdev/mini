package main

import (
	"fmt"
	"net/http"
)

type Context struct {
	Request  *Request
	Response *Response
}

func (c *Context) OkResponse(data interface{}, token string) {

	if token != "" {
		c.Response.SetHeader("Authorization ", fmt.Sprintf("Bearer %s", token))
	}
	c.Response.StatusCode(http.StatusOK).JSON(data)
	return
}

func (c *Context) CreatedResponse(data interface{}, token string) {

	if token != "" {
		c.Response.SetHeader("Authorization ", fmt.Sprintf("Bearer %s", token))
	}

	c.Response.StatusCode(http.StatusCreated).JSON(data)
	return
}

func (c *Context) InternalServerError(data interface{}, token string) {
	if token != "" {
		c.Response.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	c.Response.StatusCode(http.StatusInternalServerError).JSON(data)

	return
}

func (c *Context) BadRequest(data interface{}, token string) {
	if token != "" {
		c.Response.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	c.Response.StatusCode(http.StatusBadRequest).JSON(data)
	return
}

func (c *Context) NotFound(data interface{}, token string) {
	if token != "" {
		c.Response.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	c.Response.StatusCode(http.StatusNotFound).JSON(data)
	return
}

func (c *Context) MovedPermenantly(location, token string) {
	if token != "" {
		c.Response.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	c.Response.SetHeader("Location", location)
	c.Response.StatusCode(http.StatusMovedPermanently)
	return
}

func (c *Context) UnAuthorized(data interface{}, token string) {
	if token != "" {
		c.Response.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	c.Response.StatusCode(http.StatusUnauthorized).JSON(data)
	return
}
