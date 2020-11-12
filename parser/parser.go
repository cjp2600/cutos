package parser

import (
	"errors"
	"github.com/cjp2600/cutos/wizard"
	"github.com/getkin/kin-openapi/openapi3"
	"net/url"
)

type Type string

var CurlType Type = "curl"
var FetchType Type = "fetch"

// Path structure that we return after processing
type Path struct {
	Method         string
	SourceURL      string
	URL            *url.URL
	TemplatePath   string
	Headers        map[string]string
	PathVariables  []*Variable
	QueryVariables []*Variable
	UniqueName     string
	SourceRequest  string
	ParseRequest   map[string]interface{}
}

// HasAuthorization has auth check
func (path *Path) HasAuthorization() bool {
	if _, ok := path.Headers["authorization"]; ok {
		return true
	}
	return false
}

// Variable structure for determining path or query variables
type Variable struct {
	name        string
	varType     string
	description string
	Example     string
}

// Parser behavioral abstraction
type Parser interface {
	SetSource(source *wizard.Request)
	SetSwagger(sw *openapi3.Swagger)
	RequiredField() func(input string) error
	BuildPathMethod() *openapi3.Swagger
}

// PathParser behavioral abstraction
type PathParser interface {
	Parse() (*Path, error)
	SetText(text string)
	RequiredField() func(input string) error
}

// NewParser factory abstraction
func NewParser(t Type) (Parser, error) {
	switch t {
	case CurlType:
		return NewBuilder(NewCurl()), nil
	case FetchType:
		// todo not impl
		return NewBuilder(NewFetch()), nil
	}
	return nil, errors.New("parser type not found")
}
