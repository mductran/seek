package queries

import (
	"regexp"
	"strings"
)

type splitter struct {
	chunks []string
}

type Splitter interface {
	Chunks() []string
}

func (s *splitter) Chunks() []string {
	return s.chunks
}

func NewSplitter(sql string) Splitter {
	r := regexp.MustCompile(`\s+`)
	sql = r.ReplaceAllString(sql, " ")
	s := strings.Split(sql, " ")
	return &splitter{chunks: s}
}
