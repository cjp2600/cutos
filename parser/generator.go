package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// ConvertToSchema schematization method
func ConvertToSchema(source interface{}, rootKey string) map[string]*openapi3.SchemaRef {
	switch v := source.(type) {
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		mp := make(map[string]*openapi3.SchemaRef)
		mp[rootKey] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:    "integer",
				Format:  "int64",
				Title:   rootKey,
				Example: v,
			},
		}
		return mp
	case float64, float32:
		mp := make(map[string]*openapi3.SchemaRef)
		mp[rootKey] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:    "number",
				Title:   rootKey,
				Example: v,
			},
		}
		return mp
	case bool:
		mp := make(map[string]*openapi3.SchemaRef)
		mp[rootKey] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:    "boolean",
				Title:   rootKey,
				Example: v,
			},
		}
		return mp
	case []interface{}:
		mp := make(map[string]*openapi3.SchemaRef)
		var sub *openapi3.SchemaRef
		if len(v) > 0 {
			switch s := v[0].(type) {
			case uint8:
				sub = &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:    "string",
						Format:  "byte",
						Title:   rootKey,
						Example: v,
					},
				}
				mp[rootKey] = sub
				return mp
			default:
				var itemRef = ConvertToSchema(s, rootKey)
				sub = &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:  "array",
						Items: itemRef[rootKey],
					},
				}
			}
		} else {
			sub = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:    "string",
					Title:   rootKey,
					Example: v,
				},
			}
			mp[rootKey] = sub
		}
		mp[rootKey] = sub
		return mp
	case map[string]interface{}:
		mp := make(map[string]*openapi3.SchemaRef)
		sub := make(map[string]*openapi3.SchemaRef)
		for key, item := range v {
			mps := ConvertToSchema(item, key)
			sub[key] = mps[key]
		}
		mp[rootKey] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:    "object",
				Title:   rootKey,
				Example: v,
				Properties: sub,
			},
		}
		return mp
	default:
		mp := make(map[string]*openapi3.SchemaRef)
		mp[rootKey] = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:    "string",
				Title:   rootKey,
				Example: v,
			},
		}
		return mp
	}
}
