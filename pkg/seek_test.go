package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct{}

func TestSeek(t *testing.T) {
	seek := New()
	_, err := seek.Run(`SELECT * FROM path:testdata/data.csv AS g WHERE g.email = 'sfiles0@infoseek.co.jp'`)
	assert.Nil(t, err)
}
