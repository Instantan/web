package generate_test

import (
	"testing"
	"time"

	"github.com/Instantan/web/internal/generate"
	"github.com/Instantan/web/internal/openapi"
)

type TestStruct struct {
	ID         int                    `json:"id"`                    // Integer
	Name       string                 `json:"name"`                  // String
	Price      float64                `json:"price,omitempty"`       // Float, optional
	IsActive   bool                   `json:"is_active"`             // Boolean
	Tags       []string               `json:"tags,omitempty"`        // Slice of strings, optional
	Attributes map[string]interface{} `json:"attributes,omitempty"`  // Map with interface{}, optional
	Nested     *NestedStruct          `json:"nested,omitempty"`      // Pointer to struct, optional
	Items      []Item                 `json:"items"`                 // Slice of structs
	AnyValue   interface{}            `json:"any_value,omitempty"`   // Interface{}, optional
	CreatedAt  *time.Time             `json:"created_at,omitempty"`  // Pointer to time, optional
	UpdatedAt  *time.Time             `json:"updated_at,omitempty"`  // Pointer to time, optional
	Matrix     [][]int                `json:"matrix,omitempty"`      // Slice of slices, optional
	EmptySlice []string               `json:"empty_slice,omitempty"` // Empty slice, optional
	Anonymous  struct {               // Anonymous struct
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	} `json:"anonymous"`
}

type NestedStruct struct {
	Description string   `json:"description"`     // String
	Count       int      `json:"count,omitempty"` // Integer, optional
	SubNested   struct { // Nested anonymous struct
		Flag bool `json:"flag"` // Boolean
	} `json:"sub_nested"`
}

type Item struct {
	SKU   string  `json:"sku"`             // String
	Qty   int     `json:"qty"`             // Integer
	Price float64 `json:"price,omitempty"` // Float, optional
}

func TestGenerateTypescriptModels(t *testing.T) {
	testData := TestStruct{
		ID:       1,
		Name:     "Sample Product",
		IsActive: true,
		Tags:     []string{"tag1", "tag2"},
		Attributes: map[string]interface{}{
			"color": "red",
			"size":  "L",
			"stock": 100,
		},
		Nested: &NestedStruct{
			Description: "Nested Description",
			Count:       42,
			SubNested: struct {
				Flag bool `json:"flag"`
			}{
				Flag: true,
			},
		},
		Items: []Item{
			{SKU: "ITEM001", Qty: 10, Price: 99.99},
			{SKU: "ITEM002", Qty: 5},
		},
		AnyValue: 12345,
		Matrix:   [][]int{{1, 2}, {3, 4}},
		Anonymous: struct {
			Field1 string `json:"field1"`
			Field2 int    `json:"field2"`
		}{
			Field1: "Value1",
			Field2: 100,
		},
	}

	paths := openapi.OrderedMap[string, openapi.PathItem]{}

	paths.Set("/api/posts", openapi.PathItem{
		Get: &openapi.Operation{
			RequestBody: &openapi.RequestBody{
				Required: true,
				Content: map[string]openapi.MediaType{
					"*/*": {
						Schema: *openapi.ValueToSchema(""),
					},
				},
			},
			Description: "Get all posts",
			Responses: openapi.Responses{
				Default: openapi.Response{
					Content: map[string]openapi.MediaType{
						"*/*": {
							Schema: openapi.Schema{
								Ref: "#/components/schemas/TestStruct",
							},
						},
					},
				},
			},
		},
	})

	data := generate.GenerateTypescriptModels(openapi.OpenAPI{
		Info: openapi.Info{
			Title: "DemoApi",
		},
		Paths: paths,
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"TestStruct": *openapi.ValueToSchema(testData),
			},
		},
	})
	t.Log(string(data))
}
