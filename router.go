package mini

import (
	"reflect"
	"runtime"
)

type Route struct {
	Path        string
	Method      string
	namedRoutes *NameRoutes
	middleware  []MiddlewareFunc
}

type NameRoutes struct {
	RouteName string
	handler   HandlerFunc
}

type HandlerFunc func(*Context) error

type MiddlewareFunc func(HandlerFunc) HandlerFunc

func (handler *HandlerFunc) Name(i interface{}) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return funcName
}
