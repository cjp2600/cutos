package parser

import (
	"fmt"
	"github.com/cjp2600/cutos/config"
	"github.com/cjp2600/cutos/log"
	"github.com/cjp2600/cutos/wizard"
	"github.com/getkin/kin-openapi/openapi3"
	jsoniter "github.com/json-iterator/go"
	"github.com/manifoldco/promptui"
	"net/http"
	"strings"
)

// Builder
type Builder struct {
	sw     *openapi3.Swagger
	source *wizard.Request
	path   PathParser
}

func NewBuilder(path PathParser) *Builder {
	return &Builder{path: path}
}

func (b *Builder) SetSwagger(sw *openapi3.Swagger) {
	b.sw = sw
}

func (b *Builder) RequiredField() func(input string) error {
	return b.path.RequiredField()
}

func (b *Builder) SetSource(source *wizard.Request) {
	b.source = source
	b.path.SetText(b.source.Request)
}

func (b *Builder) BuildPathMethod() *openapi3.Swagger {
	// we process the request and divide it into components
	path, err := b.path.Parse()

	if err != nil {
		log.Fatal(err)
	}

	// Register Default Name
	var requestApiID = b.uniqueSchemeName("request", path)
	var responseApiID = b.uniqueSchemeName("response", path)

	// initialization route path
	b.pathInitialization(path)

	// set path description
	b.sw.Paths[path.TemplatePath].Description = b.source.Description

	// set tag
	if tag := b.sw.Tags.Get(b.source.Tag); tag == nil {
		if len(b.source.Tag) > 0 {
			b.sw.Tags = append(b.sw.Tags, &openapi3.Tag{
				Name: b.source.Tag,
			})
		}
	}

	// set body ref by requestApiID
	b.registerRequestBody(requestApiID, path)

	// set response
	b.registerResponse(responseApiID, path)

	// set url path variables
	b.setUrlPathVariables(path)

	// set url header
	b.setHeaders(path)

	// set url query
	b.setQueryParams(path)

	// components
	if b.sw.Components.RequestBodies == nil {
		b.sw.Components.RequestBodies = make(map[string]*openapi3.RequestBodyRef)
	}
	if b.sw.Components.Responses == nil {
		b.sw.Components.Responses = make(map[string]*openapi3.ResponseRef)
	}

	// SchemaRef
	if b.sw.Components.Schemas == nil {
		b.sw.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	}

	if len(path.SourceRequest) > 0 {
		nsr := openapi3.NewSchemaRef("#/components/schemas/"+requestApiID, nil)
		b.sw.Components.RequestBodies[requestApiID] = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.NewContentWithJSONSchemaRef(nsr),
			},
		}
	}
	if len(b.source.Response) > 0 {
		nsr := openapi3.NewSchemaRef("#/components/schemas/"+responseApiID, nil)
		b.sw.Components.Responses[responseApiID] = &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Content: openapi3.NewContentWithJSONSchemaRef(nsr),
			},
		}
	}

	if len(path.ParseRequest) > 0 {
		node := ConvertToSchema(path.ParseRequest, requestApiID)
		b.sw.Components.Schemas[requestApiID] = node[requestApiID]
	}

	if len(b.source.Response) > 0 {
		var j map[string]interface{}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		if err := json.Unmarshal([]byte(b.source.Response), &j); err != nil {
			log.Fatal(err)
		}
		node := ConvertToSchema(j, responseApiID)
		b.sw.Components.Schemas[responseApiID] = node[responseApiID]
	}

	fmt.Printf("[%s] created: %s \n", promptui.Styler(promptui.FGGreen)("✔"), path.TemplatePath)
	return b.sw
}

func (b *Builder) setHeaders(path *Path) {
	if len(path.Headers) > 0 {
		for header, value := range path.Headers {
			if !config.IsSkippedHeader(header) {
				b.sw.Paths[path.TemplatePath].Parameters = append(b.sw.Paths[path.TemplatePath].Parameters, &openapi3.ParameterRef{
					Value: &openapi3.Parameter{
						Name:    header,
						In:      "header",
						Example: value,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type:   "string",
								Format: "string",
							},
						},
					},
				})
			}
		}
	}
}

func (b *Builder) setQueryParams(path *Path) {
	if len(path.QueryVariables) > 0 {
		for _, value := range path.QueryVariables {
			b.sw.Paths[path.TemplatePath].Parameters = append(b.sw.Paths[path.TemplatePath].Parameters, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					Name:    value.name,
					In:      "query",
					Example: value.Example,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: value.varType,
						},
					},
				},
			})
		}
	}
}

func (b *Builder) setUrlPathVariables(path *Path) {
	if len(path.PathVariables) > 0 {
		for _, vr := range path.PathVariables {
			b.sw.Paths[path.TemplatePath].Parameters = append(b.sw.Paths[path.TemplatePath].Parameters, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					Name:        vr.name,
					In:          "path",
					Description: vr.description,
					Required:    true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type:   vr.varType,
							Format: "string",
						},
					},
				},
			})
		}
	}
}

func (b *Builder) uniqueSchemeName(prefix string, path *Path) string {
	return strings.Title(prefix) + path.UniqueName
}

func (b *Builder) registerResponse(responseApiID string, path *Path) {
	if len(b.source.Response) == 0 {
		return
	}
	resp := make(map[string]*openapi3.ResponseRef)
	resp["200"] = &openapi3.ResponseRef{
		Ref: "#/components/responses/" + responseApiID,
	}
	responseOperation := &openapi3.Operation{
		Tags:      []string{b.source.Tag},
		Responses: resp,
	}
	if strings.EqualFold(path.Method, http.MethodGet) {
		if b.sw.Paths[path.TemplatePath].Get == nil {
			b.sw.Paths[path.TemplatePath].Get = responseOperation
		} else {
			b.sw.Paths[path.TemplatePath].Get.Responses = resp
		}
	}
	if strings.EqualFold(path.Method, http.MethodPost) {
		if b.sw.Paths[path.TemplatePath].Post == nil {
			b.sw.Paths[path.TemplatePath].Post = responseOperation
		} else {
			b.sw.Paths[path.TemplatePath].Post.Responses = resp
		}
	}
	if strings.EqualFold(path.Method, http.MethodDelete) {
		if b.sw.Paths[path.TemplatePath].Delete == nil {
			b.sw.Paths[path.TemplatePath].Delete = responseOperation
		} else {
			b.sw.Paths[path.TemplatePath].Delete.Responses = resp
		}
	}
	if strings.EqualFold(path.Method, http.MethodPatch) {
		if b.sw.Paths[path.TemplatePath].Patch == nil {
			b.sw.Paths[path.TemplatePath].Patch = responseOperation
		} else {
			b.sw.Paths[path.TemplatePath].Patch.Responses = resp
		}
	}
	if strings.EqualFold(path.Method, http.MethodPut) {
		if b.sw.Paths[path.TemplatePath].Put == nil {
			b.sw.Paths[path.TemplatePath].Put = responseOperation
		} else {
			b.sw.Paths[path.TemplatePath].Put.Responses = resp
		}
	}
}

func (b *Builder) registerRequestBody(requestApiID string, path *Path) {
	if len(path.ParseRequest) == 0 {
		return
	}
	rb := &openapi3.RequestBodyRef{
		Ref: "#/components/requestBodies/" + requestApiID,
	}
	requestOperation := &openapi3.Operation{
		Tags:        []string{b.source.Tag},
		RequestBody: rb,
	}
	if strings.EqualFold(path.Method, http.MethodGet) {
		if b.sw.Paths[path.TemplatePath].Get == nil {
			b.sw.Paths[path.TemplatePath].Get = requestOperation
		} else {
			b.sw.Paths[path.TemplatePath].Get.RequestBody = rb
		}
	}
	if strings.EqualFold(path.Method, http.MethodPost) {
		if b.sw.Paths[path.TemplatePath].Post == nil {
			b.sw.Paths[path.TemplatePath].Post = requestOperation
		} else {
			b.sw.Paths[path.TemplatePath].Post.RequestBody = rb
		}
	}
	if strings.EqualFold(path.Method, http.MethodDelete) {
		if b.sw.Paths[path.TemplatePath].Delete == nil {
			b.sw.Paths[path.TemplatePath].Delete = requestOperation
		} else {
			b.sw.Paths[path.TemplatePath].Delete.RequestBody = rb
		}
	}
	if strings.EqualFold(path.Method, http.MethodPatch) {
		if b.sw.Paths[path.TemplatePath].Patch == nil {
			b.sw.Paths[path.TemplatePath].Patch = requestOperation
		} else {
			b.sw.Paths[path.TemplatePath].Patch.RequestBody = rb
		}
	}
	if strings.EqualFold(path.Method, http.MethodPut) {
		if b.sw.Paths[path.TemplatePath].Put == nil {
			b.sw.Paths[path.TemplatePath].Put = requestOperation
		} else {
			b.sw.Paths[path.TemplatePath].Put.RequestBody = rb
		}
	}
}

// pathInitialization initialize the patch
func (b *Builder) pathInitialization(path *Path) {
	if b.sw.Paths == nil {
		b.sw.Paths = make(map[string]*openapi3.PathItem)
	}
	if _, ok := b.sw.Paths[path.TemplatePath]; !ok {
		b.sw.Paths[path.TemplatePath] = new(openapi3.PathItem)
	}
}
