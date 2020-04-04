package swaggerdocs

import (
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

type Api interface {
	AddParamPath(p interface{}, name, desc string) Api
	AddParamPathNested(p interface{}) Api
	AddParamQuery(p interface{}, name, desc string, required bool) Api
	AddParamQueryNested(p interface{}) Api
	AddParamForm(p interface{}, name, desc string, required bool) Api
	AddParamFormNested(p interface{}) Api
	AddParamHeader(p interface{}, name, desc string, required bool) Api
	AddParamHeaderNested(p interface{}) Api
	AddParamBody(p interface{}, name, desc string, required bool) Api
	AddParamFile(name, desc string, required bool) Api
	AddResponse(code int, desc string, schema interface{}, header interface{}) Api
	SetRequestContentType(types ...string) Api
	SetResponseContentType(types ...string) Api
	SetOperationId(id string) Api
	SetDeprecated() Api
	SetDescription(desc string) Api
	SetExternalDocs(desc, url string) Api
	SetSummary(summary string) Api
	SetSecurity(names ...string) Api
	SetSecurityWithScope(s map[string][]string) Api
	Register(swagger echoswagger.Api)
}

type api struct {
	items []func(swagger echoswagger.Api)
}

func New() Api {
	return api{items: []func(swagger echoswagger.Api){}}
}

func (a api) AddParamPath(p interface{}, name, desc string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamPath(p, name, desc)
	})
	return a
}

func (a api) AddParamPathNested(p interface{}) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamPathNested(p)
	})
	return a
}

func (a api) AddParamQuery(p interface{}, name, desc string, required bool) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamQuery(p, name, desc, required)
	})
	return a
}

func (a api) AddParamQueryNested(p interface{}) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamQueryNested(p)
	})
	return a
}

func (a api) AddParamForm(p interface{}, name, desc string, required bool) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamForm(p, name, desc, required)
	})
	return a
}

func (a api) AddParamFormNested(p interface{}) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamFormNested(p)
	})
	return a
}

func (a api) AddParamHeader(p interface{}, name, desc string, required bool) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamHeader(p, name, desc, required)
	})
	return a
}

func (a api) AddParamHeaderNested(p interface{}) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamHeaderNested(p)
	})
	return a
}

func (a api) AddParamBody(p interface{}, name, desc string, required bool) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamBody(p, name, desc, required)
	})
	return a
}

func (a api) AddParamFile(name, desc string, required bool) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddParamFile(name, desc, required)
	})
	return a
}

func (a api) AddResponse(code int, desc string, schema interface{}, header interface{}) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.AddResponse(code, desc, schema, header)
	})
	return a
}

func (a api) SetRequestContentType(types ...string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetRequestContentType(types...)
	})
	return a
}

func (a api) SetResponseContentType(types ...string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetResponseContentType(types...)
	})
	return a
}

func (a api) SetOperationId(id string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetOperationId(id)
	})
	return a
}

func (a api) SetDeprecated() Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetDeprecated()
	})
	return a
}

func (a api) SetDescription(desc string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetDescription(desc)
	})
	return a
}

func (a api) SetExternalDocs(desc, url string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetExternalDocs(desc, url)
	})
	return a
}

func (a api) SetSummary(summary string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetSummary(summary)
	})
	return a
}

func (a api) SetSecurity(names ...string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetSecurity(names...)
	})
	return a
}

func (a api) SetSecurityWithScope(s map[string][]string) Api {
	a.items = append(a.items, func (swagger echoswagger.Api) {
		swagger.SetSecurityWithScope(s)
	})
	return a
}

func (a api) Route() *echo.Route {
	return  nil
}

func (a api) Register(swagger echoswagger.Api) {
	for _, f := range a.items {
		f(swagger)
	}
}
