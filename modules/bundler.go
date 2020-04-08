package modules

import (
	"dragoncore/lib/models/module"
	"dragoncore/modules/auth/oauth2"
	"dragoncore/modules/developer/content/docs"
)

var Bundler = module.Bundler().
	Bundle(developer_content_docs.Module).
	Bundle(auth_oauth2.Module)