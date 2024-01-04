package configmap

func (c *Configmap) GetSliceMapSI(key string) []map[string]interface{} {
	sliceMapSI := make([]map[string]interface{}, 0)
	sliceMapSIIntf := c.Get(key)
	if sliceMapSIIntf != nil {
		switch sliceMapSIIntf := sliceMapSIIntf.(type) {
		case []map[string]interface{}:
			sliceMapSI = sliceMapSIIntf
		case *[]map[string]interface{}:
			sliceMapSIPtr := sliceMapSIIntf
			sliceMapSI = *sliceMapSIPtr
		}
	}
	return sliceMapSI
}

func (c *Configmap) GetSlice(key string) []interface{} {
	sliceIntf := make([]interface{}, 0)
	sliceIntfIntf := c.Get(key)
	if sliceIntfIntf != nil {
		switch sliceIntfIntf := sliceIntfIntf.(type) {
		case []interface{}:
			sliceIntf = sliceIntfIntf
		}
	}
	return sliceIntf
}

func (c *Configmap) GetSliceString(key string) []string {
	sliceString := make([]string, 0)
	sliceStringIntf := c.Get(key)
	if sliceStringIntf != nil {
		switch sliceStringIntf := sliceStringIntf.(type) {
		case []string:
			sliceString = sliceStringIntf
		case *[]string:
			slicePtr := sliceStringIntf
			sliceString = *slicePtr
		}
	}
	return sliceString
}

func (c *Configmap) GetSliceByte(key string) []byte {
	sliceByte := make([]byte, 0)
	sliceByteIntf := c.Get(key)
	if sliceByteIntf != nil {
		switch sliceByteIntf := sliceByteIntf.(type) {
		case []byte:
			sliceByte = sliceByteIntf
		case *[]byte:
			slicePtr := sliceByteIntf
			sliceByte = *slicePtr
		}
	}
	return sliceByte
}
