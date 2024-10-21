package openapi

import (
	"fmt"
	"reflect"
	"strings"
)

type Schema struct {
	Type       string             `json:"type"`
	Required   []string           `json:"required,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	Items      *Schema            `json:"items,omitempty"`
	Example    any                `json:"example,omitempty"`
	TypeName   string             `json:"-"`
	// Reference to schema, if its set the the schema wont get displayed directly
	Ref string `json:"$ref,omitempty"`
}

func ValueToSchema(value any) *Schema {
	return generateSchema(value)
}

func valuesToSchemas(values []any) []Schema {
	schemas := []Schema{}
	for _, value := range values {
		schemas = append(schemas, *ValueToSchema(value))
	}
	return schemas
}

func generateSchema(value any) *Schema {
	if value == nil {
		return &Schema{Type: "null"}
	}

	s := schemaForWellKnownTypes(value)
	if s != nil {
		return s
	}

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.Bool:
		return &Schema{Type: "boolean", Example: value}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &Schema{Type: "integer", Example: value}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &Schema{Type: "integer", Example: value}
	case reflect.Float32, reflect.Float64:
		return &Schema{Type: "number", Example: value}
	case reflect.String:
		return &Schema{Type: "string", Example: value}
	case reflect.Slice, reflect.Array:
		var itemsSchema *Schema
		length := v.Len()
		if length > 0 {
			itemsSchema = generateSchema(v.Index(0).Interface())
		} else {
			elemType := t.Elem()
			zeroValue := reflect.Zero(elemType).Interface()
			itemsSchema = generateSchema(zeroValue)
		}
		return &Schema{Type: "array", Items: itemsSchema, Example: value}
	case reflect.Map:
		schema := &Schema{Type: "object", Properties: make(map[string]*Schema), Example: value}
		for _, key := range v.MapKeys() {
			propName := fmt.Sprintf("%v", key.Interface())
			propValue := v.MapIndex(key).Interface()
			schema.Properties[propName] = generateSchema(propValue)
		}
		return schema
	case reflect.Struct:
		schema := &Schema{Type: "object", Properties: make(map[string]*Schema), Example: value, TypeName: t.Name()}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			if !fieldValue.CanInterface() {
				continue
			}

			fieldName := field.Name
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				fieldName = strings.Split(jsonTag, ",")[0]
				if fieldName == "-" {
					continue
				}
			}

			fieldInterface := fieldValue.Interface()
			schema.Properties[fieldName] = generateSchema(fieldInterface)

			if !strings.Contains(jsonTag, "omitempty") && fieldName != "" {
				schema.Required = append(schema.Required, fieldName)
			}
		}
		return schema
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return &Schema{Type: "null", Example: nil}
		}
		return generateSchema(v.Elem().Interface())
	default:
		return &Schema{Type: "string", Example: fmt.Sprintf("%v", value)}
	}
}

func (o Operation) schemaOf(in string) Schema {
	s := Schema{
		Type:       "object",
		Properties: map[string]*Schema{},
		Required:   []string{},
	}
	for _, param := range o.Parameters {
		if param.In != in {
			continue
		}
		if param.Required {
			s.Required = append(s.Required, param.Name)
		}
		s.Properties[param.Name] = &param.Schema
	}
	return s
}

func (o Operation) PathSchema() Schema {
	return o.schemaOf("path")
}

func (o Operation) HeaderSchema() Schema {
	return o.schemaOf("header")
}

func (o Operation) QuerySchema() Schema {
	return o.schemaOf("query")
}

func (o Operation) CookieSchema() Schema {
	return o.schemaOf("cookie")
}
