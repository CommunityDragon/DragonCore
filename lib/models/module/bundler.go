package module

import (
	"fmt"
	"github.com/pangpanglabs/echoswagger/v2"
	"sync"
)

type bundler struct {
	modules []*module
}

// Bundler returns a bundler
func Bundler() *bundler {
	return &bundler{[]*module{}}
}

// SetInit sets the initializer
func (b *bundler) Bundle(module *module) *bundler {
	b.modules = append(b.modules, module)
	return b
}

// Register registers all bundles
func (b *bundler) Register(e echoswagger.ApiRoot) (errs []error) {
	var wg sync.WaitGroup
	for _, module := range b.modules {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := module.Register(e); err != nil {
				errs = append(errs, err)
				fmt.Printf("failed to register module %s:\n%s\n", module.name, err.Error())
			} else {
				fmt.Printf("successfully registered module %s\n", module.name)
			}
		}()
	}
	wg.Wait()
	return errs
}
