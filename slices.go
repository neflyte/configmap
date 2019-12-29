package configmap

func (c *configMap) GetSliceMapSI(key string) []map[string]interface{} {
	sliceMapSI := make([]map[string]interface{}, 0)
	sliceMapSIIntf := c.Get(key)
	if sliceMapSIIntf != nil {
		switch sliceMapSIIntf.(type) {
		case []map[string]interface{}:
			sliceMapSI = sliceMapSIIntf.([]map[string]interface{})
		}
	}
	return sliceMapSI
}

func (c *configMap) GetSlice(key string) []interface{} {
	sliceIntf := make([]interface{}, 0)
	sliceIntfIntf := c.Get(key)
	if sliceIntfIntf != nil {
		switch sliceIntfIntf.(type) {
		case []interface{}:
			sliceIntf = sliceIntfIntf.([]interface{})
		}
	}
	return sliceIntf
}
