package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionTestDb(t *testing.T) {
	asserts := assert.New(t)
	db := InitTestDb()
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")
	db.Close()
}