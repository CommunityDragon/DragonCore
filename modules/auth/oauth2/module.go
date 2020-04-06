package auth_oauth2

import (
	"dragonback/lib/models/module"
)

var Module = module.New("OAuth2 Authentication").
	SetController(handler)