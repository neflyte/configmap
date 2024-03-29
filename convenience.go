package configmap

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// AsMapSI returns the ConfigMap as its underlying data type
func (c *Configmap) AsMapSI() map[string]interface{} {
	return *c
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
			stringVal = val.String()
		}
		result[key] = stringVal
	}
	return result
}

func (c *Configmap) AsConfigmap() Configmap {
	return *c
}
