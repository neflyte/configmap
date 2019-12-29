package configmap

import "reflect"

// configMap is an alias for the underlying type (map[string]interface{})
type configMap map[string]interface{}

// ConfigMap is the interface wrapper around a map[string]interface{}
type ConfigMap interface {
	// Convenience methods
	AsMapSI() map[string]interface{}
	AsMapStringString() map[string]string
	AsJSON() []byte

	// Core methods
	Set(key string, value interface{})
	Get(key string) interface{}
	GetByKey(keys []string) interface{}
	GetType(key string) reflect.Type

	// Getters
	GetByte(key string) byte
	GetBool(key string) bool
	GetInt(key string) int
	GetUInt(key string) uint
	GetUIntPtr(key string) uintptr
	GetInt8(key string) int8
	GetUInt8(key string) uint8
	GetInt16(key string) int16
	GetUInt16(key string) uint16
	GetInt32(key string) int32
	GetUInt32(key string) uint32
	GetInt64(key string) int64
	GetUInt64(key string) uint64
	GetFloat32(key string) float32
	GetFloat64(key string) float64
	GetComplex64(key string) complex64
	GetComplex128(key string) complex128
	GetString(key string) string
	GetMapSI(key string) map[string]interface{}
	GetConfigMap(key string) ConfigMap

	// Slices
	GetSliceMapSI(key string) []map[string]interface{}
	GetSlice(key string) []interface{}

	// OrNil
	GetStringOrNil(key string) *string
}

// NewConfigMap creates a new ConfigMap, optionally based on an existing ConfigMap or map[string]interface{}
func NewConfigMap(args ...interface{}) ConfigMap {
	cm := make(configMap)
	for _, arg := range args {
		switch arg.(type) {
		case ConfigMap, configMap:
			cm = arg.(configMap)
		case *ConfigMap, *configMap:
			cmPtr := arg.(*configMap)
			cm = *cmPtr
		case map[string]interface{}:
			cm = arg.(map[string]interface{})
		case *map[string]interface{}:
			cmPtr := arg.(*map[string]interface{})
			cm = *cmPtr
		}
	}
	return &cm
}

// IsConfigMap determines if the argument is a ConfigMap or map[string]interface{}
func IsConfigMap(arg interface{}) bool {
	if arg == nil {
		return false
	}
	switch arg.(type) {
	case ConfigMap, *ConfigMap, map[string]interface{}, *map[string]interface{}, configMap, *configMap:
		return true
	}
	return false
}

// Set sets a value in the map
func (c *configMap) Set(key string, value interface{}) {
	mapSI := map[string]interface{}(*c)
	mapSI[key] = value
}

// Get retrieves a value from the map
func (c *configMap) Get(key string) interface{} {
	mapSI := map[string]interface{}(*c)
	return mapSI[key]
}

// GetByKey retrieves a value by following a path of keys; all values except the last must be ConfigMaps
func (c *configMap) GetByKey(keys []string) interface{} {
	if len(keys) > 0 {
		if len(keys) == 1 { // last key
			return c.Get(keys[0])
		}
		if IsConfigMap(c.Get(keys[0])) {
			return c.GetConfigMap(keys[0]).GetByKey(keys[1:])
		}
	}
	return nil
}

// GetType returns the reflect.Type of the value retrieved from the map
func (c *configMap) GetType(key string) reflect.Type {
	return reflect.TypeOf(c.Get(key))
}
