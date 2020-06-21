package configmap

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// AsMapSI returns the ConfigMap as its underlying data type
func (c *Configmap) AsMapSI() map[string]interface{} {
	return *c
}

// AsJSON returns the ConfigMap as a Marshaled JSON string ([]byte)
func (c *Configmap) AsJSON() []byte {
	result, err := json.Marshal(*c)
	if err != nil {
		log.Printf("error marshaling JSON: %s", err)
	}
	return result
}

// AsYAML returns the ConfigMap as a Marshaled YAML string ([]byte)
func (c *Configmap) AsYAML() []byte {
	result, err := yaml.Marshal(*c)
	if err != nil {
		log.Printf("error marshaling YAML: %s", err)
	}
	return result
}

// AsMapStringString attempts to return the ConfigMap as a map[string]string; non-string/Stringer ConfigMap values will result in empty string values
func (c *Configmap) AsMapStringString() map[string]string {
	result := make(map[string]string)
	for key, val := range *c {
		stringVal := ""
		switch val := val.(type) {
		case string:
			stringVal = val
		case *string:
			strPtr := val
			stringVal = *strPtr
		case fmt.Stringer:
			stringer := val.(fmt.Stringer)
			stringVal = stringer.String()
		}
		result[key] = stringVal
	}
	return result
}
