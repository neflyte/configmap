package configmap

func (c *configMap) GetString(key string) string {
	str := ""
	strIntf := c.Get(key)
	if strIntf != nil {
		switch strIntf.(type) {
		case string:
			str = strIntf.(string)
		}
	}
	return str
}

func (c *configMap) GetMapSI(key string) map[string]interface{} {
	mapSI := make(map[string]interface{})
	mapSIIntf := c.Get(key)
	if mapSIIntf != nil {
		switch mapSIIntf.(type) {
		case map[string]interface{}:
			mapSI = mapSIIntf.(map[string]interface{})
		}
	}
	return mapSI
}

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

func (c *configMap) GetConfigMap(key string) ConfigMap {
	cm := NewConfigMap()
	cmIntf := c.Get(key)
	if cmIntf != nil {
		switch cmIntf.(type) {
		case ConfigMap, configMap:
			cm = cmIntf.(ConfigMap)
		case *ConfigMap, *configMap:
			cmPtr := cmIntf.(*ConfigMap)
			cm = *cmPtr
		case map[string]interface{}:
			cm = NewConfigMap(cmIntf.(map[string]interface{}))
		case *map[string]interface{}:
			cmPtr := cmIntf.(*map[string]interface{})
			cm = NewConfigMap(*cmPtr)
		}
	}
	return cm
}
