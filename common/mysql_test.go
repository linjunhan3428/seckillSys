package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysql(t *testing.T) {
	_, err := NewMysqlConn()
	assert.NoError(t, err)

}
