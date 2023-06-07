package mini

import (
	"net/http"
	"strings"
)

type RouterGroup struct {
	baseURL    string
	middleware []MiddlewareFunc
	Mini       *Mini
}

func (mini *Mini) Group(prefix string) *RouterGroup {
	prefix = strings.TrimRight(prefix, "/")

	return &RouterGroup{
		baseURL:    mini.baseURL + prefix,
		middleware: []MiddlewareFunc{},
		Mini:       mini,
	}
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	prefix = strings.TrimRight(prefix, "/")

	return &RouterGroup{
		baseURL:    group.baseURL + prefix,
		middleware: group.middleware,
		Mini:       group.Mini,
	}
}

func (g *RouterGroup) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

func (g *RouterGroup) addRoutes(method, route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	g.Mini.mutex.Lock()
	defer g.Mini.mutex.Unlock()

	namedRoutes := &NameRoutes{
		RouteName: g.baseURL + route,
		handler:   handler,
	}

	g.Mini.routes = append(g.Mini.routes, &Route{
		Path:        g.baseURL + route,
		Method:      method,
		namedRoutes: namedRoutes,
		middleware:  append(g.middleware, middleware...),
	})
}

func (rg *RouterGroup) GET(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodGet, route, handler, middleware...)
}

func (rg *RouterGroup) POST(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodPost, route, handler, middleware...)
}

func (rg *RouterGroup) PUT(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodPut, route, handler, middleware...)
}

func (rg *RouterGroup) PATCH(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodPatch, route, handler, middleware...)
}

func (rg *RouterGroup) DELETE(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodDelete, route, handler, middleware...)
}

func (rg *RouterGroup) OPTIONS(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodOptions, route, handler, middleware...)
}

func (rg *RouterGroup) HEAD(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodHead, route, handler, middleware...)
}

func (rg *RouterGroup) TRACE(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodTrace, route, handler, middleware...)
}

func (rg *RouterGroup) CONNECT(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.addRoutes(http.MethodConnect, route, handler, middleware...)
}
