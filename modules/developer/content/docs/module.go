package developer_content_docs

import (
	"dragoncore/lib/models/module"
	"dragoncore/lib/models/repo"
)

var Module = module.New("Developer Documentation Content").
	SetInit(func() (err error) {
		err = repo.New(repo.CDragonDocs).Clone()
		return err
	}).
	SetController(handler)