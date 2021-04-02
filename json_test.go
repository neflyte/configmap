package configmap

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	simpleConfigmap = Configmap{
		"aString": "foo",
		"anInt":   12345,
		"aFloat":  12.345,
		"aConfigmap": Configmap{
			"aSlice": []string{"bar"},
		},
	}
	simpleConfigmapJSON            = []byte(`{"aConfigmap":{"aSlice":["bar"]},"aFloat":12.345,"aString":"foo","anInt":12345}`)
	simpleConfigmapWithJSONNumbers Configmap
)

func init() {
	simpleConfigmapWithJSONNumbers = FromJSONWithJSONNumbers(simpleConfigmapJSON).AsConfigmap()
}

func TestUnit_AsJSON(t *testing.T) {
	expected := simpleConfigmapJSON
	actual := simpleConfigmap.AsJSON()
	require.Equal(t, expected, actual)
}

func TestUnit_GetJSONNumber(t *testing.T) {
	expected := json.Number("12345")
	actual := simpleConfigmapWithJSONNumbers.GetJSONNumber("anInt")
	require.Equal(t, expected, actual)
}

func TestUnit_GetJSONNumberOrNil(t *testing.T) {
	expected := json.Number("12.345")
	actualPtr := simpleConfigmapWithJSONNumbers.GetJSONNumberOrNil("aFloat")
	require.NotNil(t, actualPtr)
	require.Equal(t, expected, *actualPtr)
}

func TestUnit_GetJSONNumberAsInt64(t *testing.T) {
	expected := int64(12345)
	actual := simpleConfigmapWithJSONNumbers.GetJSONNumberAsInt64("anInt")
	require.Equal(t, expected, actual)
}

func TestUnit_GetJSONNumberAsInt64OrNil(t *testing.T) {
	expected := int64(12345)
	actualPtr := simpleConfigmapWithJSONNumbers.GetJSONNumberAsInt64OrNil("anInt")
	require.NotNil(t, actualPtr)
	require.Equal(t, expected, *actualPtr)
}
