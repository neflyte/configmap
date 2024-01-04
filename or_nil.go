package configmap

func (c *Configmap) GetByteOrNil(key string) *byte {
	byteIntf := c.GetOrNil(key)
	if byteIntf != nil {
		switch byteIntf := byteIntf.(type) {
		case byte:
			b := byteIntf
			return &b
		case *byte:
			return byteIntf
		}
	}
	return nil
}

func (c *Configmap) GetBoolOrNil(key string) *bool {
	boolIntf := c.GetOrNil(key)
	if boolIntf != nil {
		switch boolIntf := boolIntf.(type) {
		case bool:
			b := boolIntf
			return &b
		case *bool:
			return boolIntf
		}
	}
	return nil
}

func (c *Configmap) GetIntOrNil(key string) *int {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf := intIntf.(type) {
		case int:
			i := intIntf
			return &i
		case *int:
			return intIntf
		}
	}
	return nil
}

func (c *Configmap) GetUIntOrNil(key string) *uint {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf := uintIntf.(type) {
		case uint:
			u := uintIntf
			return &u
		case *uint:
			return uintIntf
		}
	}
	return nil
}

func (c *Configmap) GetUIntPtrOrNil(key string) *uintptr {
	uintptrIntf := c.GetOrNil(key)
	if uintptrIntf != nil {
		switch uintptrIntf := uintptrIntf.(type) {
		case uintptr:
			up := uintptrIntf
			return &up
		case *uintptr:
			return uintptrIntf
		}
	}
	return nil
}

func (c *Configmap) GetInt8OrNil(key string) *int8 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf := intIntf.(type) {
		case int8:
			i := intIntf
			return &i
		case *int8:
			return intIntf
		}
	}
	return nil
}

func (c *Configmap) GetUInt8OrNil(key string) *uint8 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf := uintIntf.(type) {
		case uint8:
			u := uintIntf
			return &u
		case *uint8:
			return uintIntf
		}
	}
	return nil
}

func (c *Configmap) GetInt16OrNil(key string) *int16 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf := intIntf.(type) {
		case int16:
			i := intIntf
			return &i
		case *int16:
			return intIntf
		}
	}
	return nil
}

func (c *Configmap) GetUInt16OrNil(key string) *uint16 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf := uintIntf.(type) {
		case uint16:
			u := uintIntf
			return &u
		case *uint16:
			return uintIntf
		}
	}
	return nil
}

func (c *Configmap) GetInt32OrNil(key string) *int32 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf := intIntf.(type) {
		case int32:
			i := intIntf
			return &i
		case *int32:
			return intIntf
		}
	}
	return nil
}

func (c *Configmap) GetUInt32OrNil(key string) *uint32 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf := uintIntf.(type) {
		case uint32:
			u := uintIntf
			return &u
		case *uint32:
			return uintIntf
		}
	}
	return nil
}

func (c *Configmap) GetInt64OrNil(key string) *int64 {
	intIntf := c.GetOrNil(key)
	if intIntf != nil {
		switch intIntf := intIntf.(type) {
		case int64:
			i := intIntf
			return &i
		case *int64:
			return intIntf
		}
	}
	return nil
}

func (c *Configmap) GetUInt64OrNil(key string) *uint64 {
	uintIntf := c.GetOrNil(key)
	if uintIntf != nil {
		switch uintIntf := uintIntf.(type) {
		case uint64:
			u := uintIntf
			return &u
		case *uint64:
			return uintIntf
		}
	}
	return nil
}

func (c *Configmap) GetFloat32OrNil(key string) *float32 {
	floatIntf := c.GetOrNil(key)
	if floatIntf != nil {
		switch floatIntf := floatIntf.(type) {
		case float32:
			f := floatIntf
			return &f
		case *float32:
			return floatIntf
		}
	}
	return nil
}

func (c *Configmap) GetFloat64OrNil(key string) *float64 {
	floatIntf := c.GetOrNil(key)
	if floatIntf != nil {
		switch floatIntf := floatIntf.(type) {
		case float64:
			f := floatIntf
			return &f
		case *float64:
			return floatIntf
		}
	}
	return nil
}

func (c *Configmap) GetComplex64OrNil(key string) *complex64 {
	cplxIntf := c.GetOrNil(key)
	if cplxIntf != nil {
		switch cplxIntf := cplxIntf.(type) {
		case complex64:
			cplx := cplxIntf
			return &cplx
		case *complex64:
			return cplxIntf
		}
	}
	return nil
}

func (c *Configmap) GetComplex128OrNil(key string) *complex128 {
	cplxIntf := c.GetOrNil(key)
	if cplxIntf != nil {
		switch cplxIntf := cplxIntf.(type) {
		case complex128:
			cplx := cplxIntf
			return &cplx
		case *complex128:
			return cplxIntf
		}
	}
	return nil
}

func (c *Configmap) GetStringOrNil(key string) *string {
	strIntf := c.GetOrNil(key)
	if strIntf != nil {
		switch strIntf := strIntf.(type) {
		case string:
			str := strIntf
			return &str
		case *string:
			return strIntf
		}
	}
	return nil
}

func (c *Configmap) GetMapSIOrNil(key string) *map[string]interface{} {
	mapSIIntf := c.GetOrNil(key)
	if mapSIIntf != nil {
		switch mapSIIntf := mapSIIntf.(type) {
		case map[string]interface{}:
			mapSI := mapSIIntf
			return &mapSI
		case *map[string]interface{}:
			return mapSIIntf
		}
	}
	return nil
}

func (c *Configmap) GetConfigMapOrNil(key string) *ConfigMap {
	cmIntf := c.GetOrNil(key)
	if cmIntf != nil {
		switch cmIntf := cmIntf.(type) {
		case map[string]interface{}:
			cm := New(cmIntf)
			return &cm
		case *map[string]interface{}:
			cm := New(cmIntf)
			return &cm
		case ConfigMap:
			cm := cmIntf
			return &cm
		case *ConfigMap:
			return cmIntf
		case Configmap:
			newCM := New(cmIntf)
			return &newCM
		case *Configmap:
			cm := New(cmIntf)
			return &cm
		}
	}
	return nil
}
