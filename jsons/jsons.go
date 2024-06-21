package jsons

import (
	"encoding/json"
	"reflect"
)

func JsonEqual(a, b string) bool {
	var json1 map[string]interface{}
	var json2 map[string]interface{}

	_ = json.Unmarshal([]byte(a), &json1)
	_ = json.Unmarshal([]byte(b), &json2)

	return reflect.DeepEqual(json1, json2)
}
