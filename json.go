package configmap

import (
	"encoding/json"
	"log"
)

// AsJSON returns the ConfigMap as a Marshaled JSON string ([]byte)
func (c *Configmap) AsJSON() []byte {
	result, err := json.Marshal(*c)
	if err != nil {
		log.Printf("error marshaling JSON: %s", err)
	}
	return result
}

func (c *Configmap) GetJSONNumber(key string) json.Number {
	jsonNumber := json.Number("")
	numberIntf := c.Get(key)
	if numberIntf != nil {
		switch number := numberIntf.(type) {
		case json.Number:
			jsonNumber = number
		}
	}
	return jsonNumber
}

func (c *Configmap) GetJSONNumberOrNil(key string) *json.Number {
	var jsonNumber *json.Number
	numberIntf := c.Get(key)
	if numberIntf != nil {
		switch number := numberIntf.(type) {
		case json.Number:
			jsonNumber = &number
		}
	}
	return jsonNumber
}

func (c *Configmap) GetJSONNumberAsInt64(key string) (numberAsInt int64) {
	var err error

	numberAsInt = int64(0)
	jsonNumberIntf := c.Get(key)
	if jsonNumberIntf != nil {
		jsonNumber, ok := jsonNumberIntf.(json.Number)
		if !ok {
			return
		}
		numberAsInt, err = jsonNumber.Int64()
		if err != nil {
			log.Printf("error retrieving Int64 value of json.Number: %s", err.Error())
			return
		}
	}
	return
}

func (c *Configmap) GetJSONNumberAsFloat64(key string) (numberAsFloat float64) {
	var err error

	numberAsFloat = 0.0
	jsonNumberIntf := c.Get(key)
	if jsonNumberIntf != nil {
		jsonNumber, ok := jsonNumberIntf.(json.Number)
		if !ok {
			return
		}
		numberAsFloat, err = jsonNumber.Float64()
		if err != nil {
			log.Printf("error retrieving Float64 value of json.Number: %s", err.Error())
			return
		}
	}
	return
}

func (c *Configmap) GetJSONNumberAsInt64OrNil(key string) (numberAsInt *int64) {
	var err error
	var rawNumber int64

	numberAsInt = nil
	jsonNumberIntf := c.Get(key)
	if jsonNumberIntf != nil {
		jsonNumber, ok := jsonNumberIntf.(json.Number)
		if !ok {
			return
		}
		rawNumber, err = jsonNumber.Int64()
		if err != nil {
			log.Printf("error retrieving Int64 value of json.Number: %s", err.Error())
			return
		}
		numberAsInt = &rawNumber
	}
	return
}

func (c *Configmap) GetJSONNumberAsFloat64OrNil(key string) (numberAsFloat *float64) {
	var err error
	var rawNumber float64

	numberAsFloat = nil
	jsonNumberIntf := c.Get(key)
	if jsonNumberIntf != nil {
		jsonNumber, ok := jsonNumberIntf.(json.Number)
		if !ok {
			return
		}
		rawNumber, err = jsonNumber.Float64()
		if err != nil {
			log.Printf("error retrieving Float64 value of json.Number: %s", err.Error())
			return
		}
		numberAsFloat = &rawNumber
	}
	return
}
