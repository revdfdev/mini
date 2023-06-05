package main

type Route struct {
	Path       string
	Method     string
	handler    HandlerFunc
	middleware []MiddlewareFunc
}

type HandlerFunc func(*Context) error

type MiddlewareFunc func(HandlerFunc) HandlerFunc
