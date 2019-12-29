package configmap

func (c *Configmap) GetByteOrNil(key string) *byte {
	byteIntf := c.GetOrNil(key)
	if byteIntf != nil {
		switch byteIntf.(type) {
		case byte:
			b := byteIntf.(byte)
			return &b
		case *byte:
			return byteIntf.(*byte)
		}
	}
	return nil
}

func (c *Configmap) GetBoolOrNil(key string) *bool {
	boolIntf := c.GetOrNil(key)
	if boolIntf != nil {
		switch boolIntf.(type) {
		case bool:
			b := boolIntf.(bool)
			return &b
		case *bool:
			return boolIntf.(*bool)
		}
	}
	return nil
}

func (c *Configmap) GetIntOrNil(key string) *int {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int:
			i := intIntf.(int)
			return &i
		case *int:
			return intIntf.(*int)
		}
	}
	return nil
}

func (c *Configmap) GetUIntOrNil(key string) *uint {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint:
			u := uintIntf.(uint)
			return &u
		case *uint:
			return uintIntf.(*uint)
		}
	}
	return nil
}

func (c *Configmap) GetUIntPtrOrNil(key string) *uintptr {
	uintptrIntf := c.GetOrNil(key)
	if uintptrIntf != nil {
		switch uintptrIntf.(type) {
		case uintptr:
			up := uintptrIntf.(uintptr)
			return &up
		case *uintptr:
			return uintptrIntf.(*uintptr)
		}
	}
	return nil
}

func (c *Configmap) GetInt8OrNil(key string) *int8 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int8:
			i := intIntf.(int8)
			return &i
		case *int8:
			return intIntf.(*int8)
		}
	}
	return nil
}

func (c *Configmap) GetUInt8OrNil(key string) *uint8 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint8:
			u := uintIntf.(uint8)
			return &u
		case *uint8:
			return uintIntf.(*uint8)
		}
	}
	return nil
}

func (c *Configmap) GetInt16OrNil(key string) *int16 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int16:
			i := intIntf.(int16)
			return &i
		case *int16:
			return intIntf.(*int16)
		}
	}
	return nil
}

func (c *Configmap) GetUInt16OrNil(key string) *uint16 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint16:
			u := uintIntf.(uint16)
			return &u
		case *uint16:
			return uintIntf.(*uint16)
		}
	}
	return nil
}

func (c *Configmap) GetInt32OrNil(key string) *int32 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int32:
			i := intIntf.(int32)
			return &i
		case *int32:
			return intIntf.(*int32)
		}
	}
	return nil
}

func (c *Configmap) GetUInt32OrNil(key string) *uint32 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint32:
			u := uintIntf.(uint32)
			return &u
		case *uint32:
			return uintIntf.(*uint32)
		}
	}
	return nil
}

func (c *Configmap) GetInt64OrNil(key string) *int64 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf.(type) {
		case int64:
			i := intIntf.(int64)
			return &i
		case *int64:
			return intIntf.(*int64)
		}
	}
	return nil
}

func (c *Configmap) GetUInt64OrNil(key string) *uint64 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf.(type) {
		case uint64:
			u := uintIntf.(uint64)
			return &u
		case *uint64:
			return uintIntf.(*uint64)
		}
	}
	return nil
}

func (c *Configmap) GetFloat32OrNil(key string) *float32 {
	floatIntf := c.GetOrNil(key)
	if floatIntf != nil {
		switch floatIntf.(type) {
		case float32:
			f := floatIntf.(float32)
			return &f
		case *float32:
			return floatIntf.(*float32)
		}
	}
	return nil
}

func (c *Configmap) GetFloat64OrNil(key string) *float64 {
	floatIntf := c.GetOrNil(key)
	if floatIntf != nil {
		switch floatIntf.(type) {
		case float64:
			f := floatIntf.(float64)
			return &f
		case *float64:
			return floatIntf.(*float64)
		}
	}
	return nil
}

func (c *Configmap) GetComplex64OrNil(key string) *complex64 {
	cplxIntf := c.GetOrNil(key)
	if cplxIntf != nil {
		switch cplxIntf.(type) {
		case complex64:
			cplx := cplxIntf.(complex64)
			return &cplx
		case *complex64:
			return cplxIntf.(*complex64)
		}
	}
	return nil
}

func (c *Configmap) GetComplex128OrNil(key string) *complex128 {
	cplxIntf := c.GetOrNil(key)
	if cplxIntf != nil {
		switch cplxIntf.(type) {
		case complex128:
			cplx := cplxIntf.(complex128)
			return &cplx
		case *complex128:
			return cplxIntf.(*complex128)
		}
	}
	return nil
}

func (c *Configmap) GetStringOrNil(key string) *string {
	strIntf := c.GetOrNil(key)
	if strIntf != nil {
		switch strIntf.(type) {
		case string:
			str := strIntf.(string)
			return &str
		case *string:
			return strIntf.(*string)
		}
	}
	return nil
}

func (c *Configmap) GetMapSIOrNil(key string) *map[string]interface{} {
	mapSIIntf := c.GetOrNil(key)
	if mapSIIntf != nil {
		switch mapSIIntf.(type) {
		case map[string]interface{}:
			mapSI := mapSIIntf.(map[string]interface{})
			return &mapSI
		case *map[string]interface{}:
			return mapSIIntf.(*map[string]interface{})
		}
	}
	return nil
}

func (c *Configmap) GetConfigMapOrNil(key string) *ConfigMap {
	cmIntf := c.GetOrNil(key)
	if cmIntf != nil {
		switch cmIntf.(type) {
		case map[string]interface{}:
			mapsi := cmIntf.(map[string]interface{})
			cm := New(mapsi)
			return &cm
		case *map[string]interface{}:
			mapsiPtr := cmIntf.(*map[string]interface{})
			cm := New(mapsiPtr)
			return &cm
		case ConfigMap:
			cm := cmIntf.(ConfigMap)
			return &cm
		case *ConfigMap:
			return cmIntf.(*ConfigMap)
		case Configmap:
			cm := cmIntf.(Configmap)
			newCM := New(cm)
			return &newCM
		case *Configmap:
			cm := New(cmIntf.(*Configmap))
			return &cm
		}
	}
	return nil
}
