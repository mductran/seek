package corrector

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"seeker/queries"
	"seeker/queries/splitter"
	"testing"
)

func TestCorrectorIsCorrect(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE id < 10`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 0, len(errs))
}

func TestCorrectorMinChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
}

func TestCorrectorInvalidSelectChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidSelectChunk))
}

func TestCorrectorInvalidSelectAndFromChunk(t *testing.T) {
	sql := `SELECT * FOM path:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 2, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidSelectChunk))
	assert.True(t, errors.Is(errs[1], queries.InvalidFromChunk))
}

func TestCorrectorInvalidPathChunk(t *testing.T) {
	sql := `SELECT * FROM pth:../../testdata/data.csv AS g`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidFilePathChunk))
}

func TestCorrectorInvalidFileNotExists(t *testing.T) {
	sql := `SELECT * FROM path:../../data.csv AS g`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidFilePath))
}

func TestCorrectorInvalidAsChunk(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv ASS g`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidFromChunk))
}

func TestCorrectorWhereClause(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a b`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidWhereClause))
}

func TestCorrectorWhereClauseOperator(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a & 'b'`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidWhereClause))
}

func TestCorrectorWhereClauseValue(t *testing.T) {
	sql := `SELECT * FROM path:../../testdata/data.csv AS g WHERE a = 'b'`
	errs := IsShallowSyntaxCorrect(splitter.NewSplitter(sql))
	assert.Equal(t, 1, len(errs))
	assert.True(t, errors.Is(errs[0], queries.InvalidValueChunk))
}
