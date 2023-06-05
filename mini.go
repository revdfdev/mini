package mini

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/revdfdev/mini/utils"
)

const defaultMaxMemory = 25 * 60

type Mini struct {
	router       *http.ServeMux
	routes       []*Route
	baseURL      string
	middlewares  []MiddlewareFunc
	corsConfig   *CorsConfig
	corConfigSet bool
	mutex        sync.Mutex
}

func NewMini() *Mini {
	return &Mini{
		router:  http.NewServeMux(),
		baseURL: "",
	}
}

func (mini *Mini) addRoute(method, path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.mutex.Lock()
	defer mini.mutex.Unlock()

	mini.routes = append(mini.routes, &Route{
		Path:       mini.baseURL + path,
		Method:     method,
		handler:    handler,
		middleware: append(middleware, middleware...),
	})
}

func (mini *Mini) Use(middleware ...MiddlewareFunc) {
	mini.middlewares = append(mini.middlewares, middleware...)
}

func (mini *Mini) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := &Request{
		r,
	}

	response := &Response{w}

	c := &Context{
		Request:  request,
		Response: response,
	}

	mini.mutex.Lock()
	routes := make([]*Route, len(mini.routes))
	copy(routes, mini.routes)
	mini.mutex.Unlock()

	for _, route := range routes {
		if route.Method == r.Method && strings.HasPrefix(r.URL.Path, route.Path) {
			handler := route.handler
			for i := len(route.middleware) - 1; i >= 0; i-- {
				handler = route.middleware[i](handler)
			}
			err := handler(c)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			utils.LogResponse(r.Method, r.URL.Path, r.Response.StatusCode)

			return
		}
	}

	http.NotFound(w, r)
}

func (mini *Mini) GET(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodGet, route, handler, middleware...)
}

func (mini *Mini) POST(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodPost, route, handler, middleware...)
}

func (mini *Mini) PATCH(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodPatch, route, handler, middleware...)
}

func (mini *Mini) DELETE(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodDelete, route, handler, middleware...)
}

func (mini *Mini) OPTIONS(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodOptions, route, handler, middleware...)
}

func (mini *Mini) PUT(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodPut, route, handler, middleware...)
}

func (mini *Mini) CONNECT(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodConnect, route, handler, middleware...)
}

func (mini *Mini) TRACE(route string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	mini.addRoute(http.MethodTrace, route, handler, middleware...)
}

func (mini *Mini) Run(addr string) error {
	mini.printWelcomeMessage(addr)
	mini.printCorsStatus()
	mini.printRoutes()
	return http.ListenAndServe(addr, mini)
}

func (mini *Mini) printCorsStatus() {
	if !mini.corConfigSet {
		color.Magenta("Warning: CORS configuration not provided. CORS headers will not be added to responses.")
	}
}

func (mini *Mini) printWelcomeMessage(addr string) {
	color.Cyan("=======================================")
	color.Cyan("      Welcome to Mini REST API")
	color.Cyan("=======================================")
	fmt.Printf("Server is running on: %s\n", addr)
}

func (mini *Mini) printRoutes() {
	color.Yellow("Registered Routes:")
	for _, route := range mini.routes {
		color.Green("%s\t%s", route.Method, route.Path)
	}
	fmt.Println()
}
