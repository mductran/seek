package splitter

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSplitterWithValidSpaces(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE g.id < 10`
	s := NewSplitter(sql)

	assert.Equal(t, 9, len(s.Chunks()))
	assert.Equal(t, sql, strings.Join(s.Chunks(), " "))
}

func TestSplitterWithoutValidSpaces(t *testing.T) {
	sql := `SELECT 		* 		FROM 	path:../../testdata/data.csv AS g		WHERE 		g.id < 10`
	s := NewSplitter(sql)
	assert.Equal(t, 9, len(s.Chunks()))
}
