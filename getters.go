package configmap

func (c *Configmap) GetByte(key string) byte {
	b := byte(0)
	byteIntf := c.Get(key)
	if byteIntf != nil {
		switch byteIntf.(type) {
		case byte:
			b = byteIntf.(byte)
		}
	}
	return b
}

func (c *Configmap) GetBool(key string) bool {
	b := false
	boolIntf := c.Get(key)
	if boolIntf != nil {
		switch boolIntf.(type) {
		case bool:
			b = boolIntf.(bool)
		}
	}
	return b
}

func (c *Configmap) GetInt(key string) int {
	i := 0
	intIntf := c.Get(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int:
			i = intIntf.(int)
		}
	}
	return i
}

func (c *Configmap) GetUInt(key string) uint {
	u := uint(0)
	uintIntf := c.Get(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint:
			u = uintIntf.(uint)
		}
	}
	return u
}

func (c *Configmap) GetUIntPtr(key string) uintptr {
	up := uintptr(0)
	uintptrIntf := c.Get(key)
	if uintptrIntf != nil {
		switch uintptrIntf.(type) {
		case uintptr:
			up = uintptrIntf.(uintptr)
		}
	}
	return up
}

func (c *Configmap) GetInt8(key string) int8 {
	i := int8(0)
	intIntf := c.Get(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int8:
			i = intIntf.(int8)
		}
	}
	return i
}

func (c *Configmap) GetUInt8(key string) uint8 {
	u := uint8(0)
	uintIntf := c.Get(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint8:
			u = uintIntf.(uint8)
		}
	}
	return u
}

func (c *Configmap) GetInt16(key string) int16 {
	i := int16(0)
	intIntf := c.Get(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int16:
			i = intIntf.(int16)
		}
	}
	return i
}

func (c *Configmap) GetUInt16(key string) uint16 {
	u := uint16(0)
	uintIntf := c.Get(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint16:
			u = uintIntf.(uint16)
		}
	}
	return u
}

func (c *Configmap) GetInt32(key string) int32 {
	i := int32(0)
	intIntf := c.Get(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int32:
			i = intIntf.(int32)
		}
	}
	return i
}

func (c *Configmap) GetUInt32(key string) uint32 {
	u := uint32(0)
	uintIntf := c.Get(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint32:
			u = uintIntf.(uint32)
		}
	}
	return u
}

func (c *Configmap) GetInt64(key string) int64 {
	i := int64(0)
	intIntf := c.Get(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int64:
			i = intIntf.(int64)
		}
	}
	return i
}

func (c *Configmap) GetUInt64(key string) uint64 {
	u := uint64(0)
	uintIntf := c.Get(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint64:
			u = uintIntf.(uint64)
		}
	}
	return u
}

func (c *Configmap) GetFloat32(key string) float32 {
	f := float32(0.0)
	floatIntf := c.Get(key)
	if floatIntf != nil {
		switch floatIntf.(type) {
		case float32:
			f = floatIntf.(float32)
		}
	}
	return f
}

func (c *Configmap) GetFloat64(key string) float64 {
	f := 0.0
	floatIntf := c.Get(key)
	if floatIntf != nil {
		switch floatIntf.(type) {
		case float64:
			f = floatIntf.(float64)
		}
	}
	return f
}

func (c *Configmap) GetComplex64(key string) complex64 {
	cplx := complex64(0)
	cplxIntf := c.Get(key)
	if cplxIntf != nil {
		switch cplxIntf.(type) {
		case complex64:
			cplx = cplxIntf.(complex64)
		}
	}
	return cplx
}

func (c *Configmap) GetComplex128(key string) complex128 {
	cplx := complex128(0)
	cplxIntf := c.Get(key)
	if cplxIntf != nil {
		switch cplxIntf.(type) {
		case complex128:
			cplx = cplxIntf.(complex128)
		}
	}
	return cplx
}

func (c *Configmap) GetString(key string) string {
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

func (c *Configmap) GetMapSI(key string) map[string]interface{} {
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

func (c *Configmap) GetConfigMap(key string) ConfigMap {
	cm := NewConfigMap()
	cmIntf := c.Get(key)
	if cmIntf != nil {
		switch cmIntf.(type) {
		case ConfigMap, Configmap:
			cm = cmIntf.(ConfigMap)
		case *ConfigMap, *Configmap:
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
