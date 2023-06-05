package mini

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Context struct {
	Request  *Request
	Response *Response
	Keys     map[string]any
	mu       sync.RWMutex
}

func (c *Context) BindJson(obj interface{}) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if err := json.NewDecoder(c.Request.Body).Decode(obj); err != nil {
		return err
	}

	return nil
}

func (c *Context) ParseMultiParForm() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if err := c.Request.ParseMultipartForm(defaultMaxMemory); err != nil {
		return err
	}

	return nil
}

func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.Keys[key]
	return
}

func (c *Context) ShouldGet(key string) any {
	if value, exists := c.Get(key); exists {
		return value
	}

	return nil
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
