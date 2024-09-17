package db

import (
	"fmt"
	"seeker/pkg/queries"
)

type db struct {
	structure queries.Structure
}

type DB interface {
	Run() (map[string]string, error)
}

func (d *db) Run() (map[string]string, error) {
	file := d.structure.FileDB()
	fmt.Println(file.Alias(), file.Path())

	return nil, nil
}

func NewDb(s queries.Structure) DB {
	return &db{structure: s}
}
