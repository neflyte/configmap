package configmap

func (c *Configmap) GetSliceMapSI(key string) []map[string]interface{} {
	sliceMapSI := make([]map[string]interface{}, 0)
	sliceMapSIIntf := c.Get(key)
	if sliceMapSIIntf != nil {
		switch sliceMapSIIntf.(type) {
		case []map[string]interface{}:
			sliceMapSI = sliceMapSIIntf.([]map[string]interface{})
		case *[]map[string]interface{}:
			sliceMapSIPtr := sliceMapSIIntf.(*[]map[string]interface{})
			sliceMapSI = *sliceMapSIPtr
		}
	}
	return sliceMapSI
}

func (c *Configmap) GetSlice(key string) []interface{} {
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

func (c *Configmap) GetSliceString(key string) []string {
	sliceString := make([]string, 0)
	sliceStringIntf := c.Get(key)
	if sliceStringIntf != nil {
		switch sliceStringIntf.(type) {
		case []string:
			sliceString = sliceStringIntf.([]string)
		case *[]string:
			slicePtr := sliceStringIntf.(*[]string)
			sliceString = *slicePtr
		}
	}
	return sliceString
}

func (c *Configmap) GetSliceByte(key string) []byte {
	sliceByte := make([]byte, 0)
	sliceByteIntf := c.Get(key)
	if sliceByteIntf != nil {
		switch sliceByteIntf.(type) {
		case []byte:
			sliceByte = sliceByteIntf.([]byte)
		case *[]byte:
			slicePtr := sliceByteIntf.(*[]byte)
			sliceByte = *slicePtr
		}
	}
	return sliceByte
}
