package queries

import (
	"github.com/stretchr/testify/assert"
	"seeker/queries/syntax"
	"testing"
)

func TestStructureValid(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE g.id < 10`
	res := NewStructure(sql)

	assert.Equal(t, false, res.HasErrors())
	assert.Equal(t, res.Result().Column().Type(), syntax.ColumnType)
	assert.Equal(t, res.Result().FileDB().Type(), syntax.FileDbType)
}
