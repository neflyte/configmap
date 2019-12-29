package configmap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMap = &Configmap{
		"int":   1,
		"int8":  int8(2),
		"int16": int16(3),
		"int32": int32(4),
		"int64": int64(5),
		"map": &Configmap{
			"key": "value",
		},
	}
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

func TestUnit_StructLiteral(t *testing.T) {
	foo := &Configmap{"key": "value"}
	assert.Equal(t, "value", foo.GetString("key"))
}

func TestUnit_TestMap(t *testing.T) {
	assert.Equal(t, 1, testMap.GetInt("int"))
	assert.Equal(t, int8(2), testMap.GetInt8("int8"))
	assert.Equal(t, int16(3), testMap.GetInt16("int16"))
	assert.Equal(t, int32(4), testMap.GetInt32("int32"))
	assert.Equal(t, int64(5), testMap.GetInt64("int64"))
	assert.Equal(t, &Configmap{"key": "value"}, testMap.GetConfigMap("map"))
}
