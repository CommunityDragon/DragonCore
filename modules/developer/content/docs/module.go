package developer_content_docs

import (
	"dragonback/lib/models/module"
	"dragonback/lib/models/repo"
)

var Module = module.New("Developer Documentation").
	SetInit(func() (err error) {
		err = repo.New(repo.CDragonDocs).Clone()
		return err
	}).
	SetController(handler)