package queries

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorrectorIsCorrect(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE id < 10`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 0, len(errs))
}

func TestCorrectorMinChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
}

func TestCorrectorInvalidSelectChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidSelectChunk))
}

func TestCorrectorInvalidSelectAndFromChunk(t *testing.T) {
	sql := `SELECT * FOM path:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 2, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidSelectChunk))
	assert.True(t, errors.Is(errs[1], InvalidFromChunk))
}

func TestCorrectorInvalidPathChunk(t *testing.T) {
	sql := `SELECT * FROM pth:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidFilePathChunk))
}

func TestCorrectorINvalidFileNotExists(t *testing.T) {
	sql := `SELECT * FROM path:../../data.csv AS g`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidFilePath))
}

func TestCorrectorInvalidAsChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv ASS g`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidFromChunk))
}

func TestCorrectorWhereClause(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a b`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidWhereClause))
}

func TestCorrectorWhereClauseOperator(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a & 'b'`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidWhereClause))
}

func TestCorrectorWhereClauseValue(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a = 'b'`
	errs := IsShallowSyntaxCorrect(NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], InvalidValueChunk))
}
