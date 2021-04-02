package configmap

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"log"
	"reflect"
)

// Version is the version number of the library
const Version = "0.2.0"

// Configmap is an alias for the underlying type (map[string]interface{})
type Configmap map[string]interface{}

// ConfigMap is the interface wrapper around a map[string]interface{}
type ConfigMap interface {
	// Convenience methods
	AsMapSI() map[string]interface{}
	AsMapStringString() map[string]string
	AsYAML() []byte
	AsConfigmap() Configmap

	// Core methods
	Set(key string, value interface{})
	Get(key string) interface{}
	Has(key string) bool
	Delete(keys ...string)
	GetOrNil(key string) interface{}
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
	GetSlice(key string) []interface{}
	GetSliceMapSI(key string) []map[string]interface{}
	GetSliceString(key string) []string
	GetSliceByte(key string) []byte

	// OrNil
	GetByteOrNil(key string) *byte
	GetBoolOrNil(key string) *bool
	GetIntOrNil(key string) *int
	GetUIntOrNil(key string) *uint
	GetUIntPtrOrNil(key string) *uintptr
	GetInt8OrNil(key string) *int8
	GetUInt8OrNil(key string) *uint8
	GetInt16OrNil(key string) *int16
	GetUInt16OrNil(key string) *uint16
	GetInt32OrNil(key string) *int32
	GetUInt32OrNil(key string) *uint32
	GetInt64OrNil(key string) *int64
	GetUInt64OrNil(key string) *uint64
	GetFloat32OrNil(key string) *float32
	GetFloat64OrNil(key string) *float64
	GetComplex64OrNil(key string) *complex64
	GetComplex128OrNil(key string) *complex128
	GetStringOrNil(key string) *string
	GetMapSIOrNil(key string) *map[string]interface{}
	GetConfigMapOrNil(key string) *ConfigMap

	// JSON methods
	AsJSON() []byte
	GetJSONNumber(key string) json.Number
	GetJSONNumberOrNil(key string) *json.Number
	GetJSONNumberAsInt64(key string) (numberAsInt int64)
	GetJSONNumberAsFloat64(key string) (numberAsFloat float64)
	GetJSONNumberAsInt64OrNil(key string) (numberAsInt *int64)
	GetJSONNumberAsFloat64OrNil(key string) (numberAsFloat *float64)
}

// New creates a new ConfigMap, optionally based on an existing ConfigMap or map[string]interface{}
func New(args ...interface{}) ConfigMap {
	cm := Configmap{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case ConfigMap:
			return arg
		case *ConfigMap:
			return *arg
		case map[string]interface{}:
			cm = arg
		case *map[string]interface{}:
			cmPtr := arg
			cm = *cmPtr
		}
	}
	return &cm
}

// FromJSON returns a ConfigMap that was unmarshalled from a JSON blob
func FromJSON(rawJSON []byte) ConfigMap {
	mapsi := make(map[string]interface{})
	err := json.Unmarshal(rawJSON, &mapsi)
	if err != nil {
		log.Printf("error unmarshaling JSON: %s", err)
	}
	return New(mapsi)
}

// FromJSONWithJSONNumbers returns a ConfigMap that was unmarshalled from a JSON blob
// using json.Number for numeric fields
func FromJSONWithJSONNumbers(rawJSON []byte) ConfigMap {
	mapsi := make(map[string]interface{})
	dec := json.NewDecoder(bytes.NewReader(rawJSON))
	dec.UseNumber()
	err := dec.Decode(&mapsi)
	if err != nil {
		log.Printf("error decoding JSON: %s", err)
	}
	return New(mapsi)
}

// FromYAML returns a ConfigMap that was unmarshalled from a YAML document
func FromYAML(rawYAML []byte) ConfigMap {
	mapsi := make(map[string]interface{})
	err := yaml.Unmarshal(rawYAML, &mapsi)
	if err != nil {
		log.Printf("error unmarshaling YAML: %s", err)
	}
	return New(mapsi)
}

// IsConfigMap determines if the argument is a ConfigMap or map[string]interface{}
func IsConfigMap(arg interface{}) bool {
	if arg == nil {
		return false
	}
	switch arg.(type) {
	case ConfigMap, *ConfigMap, map[string]interface{}, *map[string]interface{}, Configmap, *Configmap:
		return true
	}
	return false
}

// Has determines if a particular key exists in the map
func (c *Configmap) Has(key string) bool {
	_, ok := map[string]interface{}(*c)[key]
	return ok
}

// Delete removes a number of keys and their values from the map
func (c *Configmap) Delete(keys ...string) {
	mapSI := map[string]interface{}(*c)
	for _, key := range keys {
		delete(mapSI, key)
	}
}

// Set sets a value in the map
func (c *Configmap) Set(key string, value interface{}) {
	mapSI := map[string]interface{}(*c)
	mapSI[key] = value
}

// Get retrieves a value from the map
func (c *Configmap) Get(key string) interface{} {
	mapSI := map[string]interface{}(*c)
	return mapSI[key]
}

// GetOrNil attempts to retrieve a value from the map; if no value corresponds to the supplied key, nii is returned
func (c *Configmap) GetOrNil(key string) interface{} {
	mapSI := map[string]interface{}(*c)
	val, ok := mapSI[key]
	if !ok {
		return nil
	}
	return val
}

// GetByKey retrieves a value by following a path of keys; all values except the last must be ConfigMaps
func (c *Configmap) GetByKey(keys []string) interface{} {
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
func (c *Configmap) GetType(key string) reflect.Type {
	return reflect.TypeOf(c.Get(key))
}
