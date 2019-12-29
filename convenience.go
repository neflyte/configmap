package configmap

import (
	"encoding/json"
	"fmt"

	"github.com/go-yaml/yaml"
)

// AsMapSI returns the ConfigMap as its underlying data type
func (c *Configmap) AsMapSI() map[string]interface{} {
	return *c
}

// AsJSON returns the ConfigMap as a Marshaled JSON string ([]byte)
func (c *Configmap) AsJSON() []byte {
	result := make([]byte, 0)
	result, _ = json.Marshal(*c)
	return result
}

// AsYAML returns the ConfigMap as a Marshaled YAML string ([]byte)
func (c *Configmap) AsYAML() []byte {
	result := make([]byte, 0)
	result, _ = yaml.Marshal(*c)
	return result
}

// AsMapStringString attempts to return the ConfigMap as a map[string]string; non-string/Stringer ConfigMap values will result in empty string values
func (c *Configmap) AsMapStringString() map[string]string {
	result := make(map[string]string)
	for key, val := range *c {
		stringVal := ""
		switch val.(type) {
		case string:
			stringVal = val.(string)
		case *string:
			strPtr := val.(*string)
			stringVal = *strPtr
		case fmt.Stringer:
			stringer := val.(fmt.Stringer)
			stringVal = stringer.String()
		}
		result[key] = stringVal
	}
	return result
}
