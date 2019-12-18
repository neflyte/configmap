package configmap

// AsMapSI returns the ConfigMap as its underlying data type
func (c *configMap) AsMapSI() map[string]interface{} {
	return *c
}
