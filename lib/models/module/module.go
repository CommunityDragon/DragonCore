package module

import (
	"dragonback/lib/models/controller"
	"github.com/pangpanglabs/echoswagger/v2"
)

type module struct {
	name string
	init func() error
	controller controller.Controller
}

func New(name string) *module {
	return &module{name, nil, nil}
}

// SetInit sets the initializer
func (m *module) SetInit(init func() error) *module {
	m.init = init
	return m
}

// SetController sets the controller
func (m *module) SetController(controller controller.Controller) *module {
	controller.Name(m.name)
	m.controller = controller
	return m
}

// Register registers the module
func (m *module) Register(e echoswagger.ApiRoot) (err error) {
	if m.init != nil {
		if err = m.init(); err != nil {
			return err
		}
	}

	if m.controller != nil {
		m.controller.Register(e)
	}

	return nil
}