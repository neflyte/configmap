package configmap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUnit_Set(t *testing.T) {
	str := "bar"
	cm := NewConfigMap()
	cm.Set("foo", str)
	act := cm.Get("foo")
	assert.Equal(t, str, act)
}

func TestUnit_GetByKey_Basic(t *testing.T) {
	expected := "fnord"
	data := NewConfigMap(map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": expected,
			},
		},
	})
	actual := data.GetByKey([]string{"foo", "bar", "baz"})
	assert.Equal(t, expected, actual)
}

func TestUnit_GetByKey_Negative(t *testing.T) {
	data := NewConfigMap(map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": "narf",
		},
	})
	actual := data.GetByKey([]string{"foo", "bar", "baz"})
	assert.Nil(t, actual)
}
