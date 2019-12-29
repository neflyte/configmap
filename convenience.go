package configmap

import "encoding/json"

// AsMapSI returns the ConfigMap as its underlying data type
func (c *configMap) AsMapSI() map[string]interface{} {
	return *c
}

// AsJSON returns the ConfigMap as a Marshaled JSON string ([]byte)
func (c *configMap) AsJSON() []byte {
	result := make([]byte, 0)
	result, _ = json.Marshal(*c)
	return result
}

// AsMapStringString attempts to return the ConfigMap as a map[string]string; non-string ConfigMap values will result in empty string values
func (c *configMap) AsMapStringString() map[string]string {
	result := make(map[string]string)
	for key, val := range *c {
		stringVal := ""
		switch val.(type) {
		case string:
			stringVal = val.(string)
		case *string:
			strPtr := val.(*string)
			stringVal = *strPtr
		}
		result[key] = stringVal
	}
	return result
}
