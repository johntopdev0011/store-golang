package jsonSerializer

import (
	"github.com/TylerBrock/colorjson"
	"github.com/goccy/go-json"
	typeMapper "github.com/mehdihadeli/store-golang-microservice-sample/pkg/reflection/type_mappper"
	"github.com/mitchellh/mapstructure"
	"log"
)

//https://www.sohamkamani.com/golang/json/#decoding-json-to-maps---unstructured-data
//https://developpaper.com/mapstructure-of-go/
//https://github.com/goccy/go-json

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal is a wrapper around json.Unmarshal.
// To unmarshal JSON into an interface value, Unmarshal stores in a map[string]interface{}
func Unmarshal(data []byte, v interface{}) error {

	//https://pkg.go.dev/encoding/json#Unmarshal
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	log.Printf("deserialize structure object")

	return nil
}

func UnmarshalT[T any](data []byte) error {
	v := typeMapper.InstanceByT[T]()
	//https://pkg.go.dev/encoding/json#Unmarshal
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	log.Printf("deserialize structure object")

	return nil
}

// UnmarshalFromJson is a wrapper around json.Unmarshal.
func UnmarshalFromJson(data string, v interface{}) error {
	err := Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}

	return nil
}

// DecodeWithMapStructure is a wrapper around mapstructure.Decode.
// Decode takes an input structure or map[string]interface{} and uses reflection to translate it to the output structure. output must be a pointer to a map or struct.
// https://pkg.go.dev/github.com/mitchellh/mapstructure#section-readme
func DecodeWithMapStructure(input interface{}, output interface{}) error {
	//https://developpaper.com/mapstructure-of-go/
	return mapstructure.Decode(input, output)
}

func UnmarshalToMap(data []byte, v *map[string]interface{}) error {
	//https://developpaper.com/mapstructure-of-go/
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalToMapFromJson(data string, v *map[string]interface{}) error {
	return UnmarshalToMap([]byte(data), v)
}

// PrettyPrint print input object as a formatted json string
func PrettyPrint(data interface{}) string {
	//https://gosamples.dev/pretty-print-json/
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

// ColoredPrettyPrint print input object as a formatted json string with color
func ColoredPrettyPrint(data interface{}) string {
	//https://github.com/TylerBrock/colorjson
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(PrettyPrint(data)), &obj)
	if err != nil {
		return ""
	}
	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	val, err := f.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(val)
}
