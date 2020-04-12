package configmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit_GetStringOrNil(t *testing.T) {
	cm := New()
	cm.Set("foo", "bar")
	baz := cm.GetStringOrNil("baz")
	bar := cm.GetString("foo")
	assert.Nil(t, baz)
	assert.Equal(t, "bar", bar)
}
