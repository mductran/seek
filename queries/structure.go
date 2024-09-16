package queries

import (
	"fmt"
	"seeker/queries/syntax"
	"strings"
)

type structure struct {
	column syntax.SyntaxType
	fileDb syntax.SyntaxType
}

type Structure interface {
	Column() syntax.SyntaxType
	FileDB() syntax.SyntaxType
}

func (s *structure) Column() syntax.SyntaxType {
	return s.column
}

func (s *structure) FileDB() syntax.SyntaxType {
	return s.fileDb
}

func NewStructure(sql string) Result[Structure] {
	s := NewSplitter(sql)
	errs := IsShallowSyntaxCorrect(s)
	if len(errs) != 0 {
		return NewResult[Structure](nil, errs)
	}

	columns := splitColumns(s.Chunks()[1])
	f, alias := resolveFiles(s.Chunks()[3], s.Chunks()[5])

	syntaxStructure := structure{
		column: syntax.NewColumn(columns),
		fileDb: syntax.NewFileDb(f, alias),
	}

	resolveWhereClause(s.Chunks()[5:])

	return NewResult[Structure](&syntaxStructure, []error{})
}

func splitColumns(c string) []string {
	if c == "*" {
		return []string{"*"}
	}
	return strings.Split(c, ",")
}

func resolveFiles(path, alias string) (string, string) {
	p := strings.Split(path, ":")
	return p[1], alias
}

func resolveWhereClause(chunks []string) {
	fmt.Println(chunks)
}
