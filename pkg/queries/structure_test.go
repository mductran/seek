package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructureValid(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE g.email = 'sfiles0@infoseek.co.jp'`
	res := NewStructure(sql)

	assert.Equal(t, false, res.HasErrors())
	assert.Nil(t, res.Errors())

	condition := res.Result().Condition()
	assert.Equal(t, condition.Column(), "email")
	assert.Equal(t, condition.Value(), "'sfiles0@infoseek.co.jp'")
	assert.Equal(t, condition.Operator(), "=")
	assert.Nil(t, condition.Next())
}
