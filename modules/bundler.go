package modules

import (
	"dragonback/lib/models/module"
	"dragonback/modules/auth/oauth2"
	"dragonback/modules/developer/content/docs"
)

var Bundler = module.Bundler().
	Bundle(developer_content_docs.Module).
	Bundle(auth_oauth2.Module)