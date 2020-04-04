package controller

import (
	"dragonback/lib/models/swaggerdocs"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"net/http"
)

type Method string

const (
	GET   		Method = http.MethodGet
	POST  		Method = http.MethodPost
	PUT    		Method = http.MethodPut
	DELETE   	Method = http.MethodDelete
	PATCH    	Method = http.MethodPatch
	OPTIONS  	Method = http.MethodOptions
	CONNECT   	Method = http.MethodConnect
	TRACE  		Method = http.MethodTrace
	HEAD 		Method = http.MethodHead
)

type Controller interface {
	Name(name string) *controller
	Middleware(mw echo.MiddlewareFunc) *controller
	Handler(method Method, path string, api swaggerdocs.Api, handle echo.HandlerFunc, mw ...echo.MiddlewareFunc) *controller
	Register(e echoswagger.ApiRoot)
}

// controller struct
type controller struct {
	name string
	prefix string
	mw []echo.MiddlewareFunc
	h []handler
}

// handler struct
type handler struct {
	method Method
	path string
	api swaggerdocs.Api
	h echo.HandlerFunc
	mw []echo.MiddlewareFunc
}

// New creates a controller
func New(prefix string) *controller {
	return &controller{"", prefix, []echo.MiddlewareFunc{}, []handler{}}
}

// Sets the controller name
func (c *controller) Name(name string) *controller {
	c.name = name
	return c
}

// Middleware adds a controller middleware
func (c *controller) Middleware(mw echo.MiddlewareFunc) *controller {
	c.mw = append(c.mw, mw)
	return c
}

// Handler adds a controller handler
func (c *controller) Handler(method Method, path string, api swaggerdocs.Api, handle echo.HandlerFunc, mw ...echo.MiddlewareFunc) *controller {
	c.h = append(c.h, handler{method, path, api, handle, mw})
	return c
}

// Register registers the controller
func (c *controller) Register(e echoswagger.ApiRoot) {
	g := e.Group(c.name, c.prefix, c.mw...)
	for _, h := range c.h {
		if api := g.Add(string(h.method), h.path, h.h, h.mw...); h.api != nil {
			h.api.Register(api)
		}
	}
}
