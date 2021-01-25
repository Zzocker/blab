package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestNewForTest(t *testing.T) {
	assert := assert.New(t)
	logger, entries := NewForTest()
	assert.Zero(entries.Len())
	logger.Info("msg 1")
	assert.Equal(1, entries.Len())
	logger.Info("msg 2")
	logger.Info("msg 3")
	assert.Equal(3, entries.Len())
	entries.TakeAll()
	assert.Zero(entries.Len())
	logger.Info("msg 4")
	assert.Equal(1, entries.Len())
}
