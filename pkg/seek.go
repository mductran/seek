package pkg

import (
	"seeker/pkg/db"
	"seeker/pkg/queries"
)

type seek struct{}

type Seek interface {
	Run(sql string) (map[string]string, []error)
}

func (s *seek) Run(sql string) (map[string]string, []error) {
	res := queries.NewStructure(sql)
	if res.HasErrors() {
		return nil, res.Errors()
	}

	dbRunner := db.NewDb(res.Result())
	_, err := dbRunner.Run()
	if err != nil {
		return nil, nil
	}

	return nil, nil
}

func New() Seek {
	return &seek{}
}
