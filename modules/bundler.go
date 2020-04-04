package modules

import (
	"dragonback/lib/models/module"
	"dragonback/modules/developer/content/docs"
)

var Bundler = module.Bundler().
	Bundle(developer_content_docs.Module)