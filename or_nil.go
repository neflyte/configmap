package configmap

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
