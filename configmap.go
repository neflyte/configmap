package configmap

type configMap map[string]interface{}

type ConfigMap interface {
	AsMapSI() map[string]interface{}
	Set(key string, value interface{})
	Get(key string) interface{}
	GetString(key string) string
	GetStringOrNil(key string) *string
	GetMapSI(key string) map[string]interface{}
	GetConfigMap(key string) ConfigMap
}

func NewConfigMap(args ...interface{}) ConfigMap {
	mapSI := make(map[string]interface{})
	cm := configMap(mapSI)
	for _, arg := range args {
		switch arg.(type) {
		case map[string]interface{}:
			cm = arg.(map[string]interface{})
		}
	}
	return &cm
}

func (c *configMap) AsMapSI() map[string]interface{} {
	return *c
}

func (c *configMap) Set(key string, value interface{}) {
	mapSI := map[string]interface{}(*c)
	mapSI[key] = value
}

func (c *configMap) Get(key string) interface{} {
	mapSI := map[string]interface{}(*c)
	return mapSI[key]
}

func (c *configMap) GetString(key string) string {
	str := ""
	strIntf := c.Get(key)
	switch strIntf.(type) {
	case string:
		str = strIntf.(string)
	}
	return str
}

func (c *configMap) GetStringOrNil(key string) *string {
	strIntf := c.Get(key)
	if strIntf == nil {
		return nil
	}
	switch strIntf.(type) {
	case string:
		str := strIntf.(string)
		return &str
	}
	return nil
}

func (c *configMap) GetMapSI(key string) map[string]interface{} {
	mapSI := make(map[string]interface{})
	mapSIIntf := c.Get(key)
	switch mapSIIntf.(type) {
	case map[string]interface{}:
		mapSI = mapSIIntf.(map[string]interface{})
	}
	return mapSI
}

func (c *configMap) GetConfigMap(key string) ConfigMap {
	cm := NewConfigMap()
	cmIntf := c.Get(key)
	switch cmIntf.(type) {
	case ConfigMap:
		cm = cmIntf.(ConfigMap)
	case map[string]interface{}:
		cm = NewConfigMap(cmIntf.(map[string]interface{}))
	}
	return cm
}
