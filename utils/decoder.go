package utils

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	jsoniter "github.com/json-iterator/go"
	"path/filepath"
	"strings"
)

type MarshalType string

var Yaml MarshalType = "yaml"
var Json MarshalType = "json"

func FindEx(fileName string) MarshalType {
	extension := filepath.Ext(fileName)
	switch strings.ToLower(extension) {
	case ".yaml", ".yml":
		return Yaml
	default:
		return Json
	}
}

func Marshal(fileName string, swagger *openapi3.Swagger) ([]byte, error) {
	switch FindEx(fileName) {
	case Yaml:
		return YamlMarshal(swagger)
	default:
		return JsonMarshal(swagger)
	}
}

func JsonMarshal(swagger *openapi3.Swagger) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(swagger)
}

func YamlMarshal(swagger *openapi3.Swagger) ([]byte, error) {
	return yaml.Marshal(swagger)
}
