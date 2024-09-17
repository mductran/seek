package db

import "seeker/pkg/queries"

type db struct {
	structure queries.Structure
}

type DB interface {
	Run() (map[string]string, error)
}

func (d *db) Run() (map[string]string, error) {
	return nil, nil
}

func NewDb(s queries.Structure) DB {
	return &db{structure: s}
}
