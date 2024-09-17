package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct{}

func TestSeek(t *testing.T) {
	seek := New()
	_, err := seek.Run(`SELECT * FROM path:testdata/data.csv AS g WHERE first_name = 'Skye'`)
	assert.Nil(t, err)
}
