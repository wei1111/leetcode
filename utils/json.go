package utils

import (
	"bytes"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func JsonStr(v interface{}) string {
	result, _ := jsoniter.Marshal(v)
	return string(result)
}

func PrettyJsonStr(v interface{}) string {
	result, _ := jsoniter.MarshalIndent(v, "", " ")
	return string(result)
}

func UnmarshalWithNumber(jsonBytes []byte, result interface{}) error {
	decoder := jsoniter.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.UseNumber()
	return decoder.Decode(result)
}

func UnmarshalWithNumberFromStr(jsonStr string, result interface{}) error {
	decoder := jsoniter.NewDecoder(strings.NewReader(jsonStr))
	decoder.UseNumber()
	return decoder.Decode(result)
}
