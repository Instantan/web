package openapi_test

import (
	"encoding/json"
	"testing"

	"github.com/Instantan/web/internal/openapi"
)

func TestValueToSchema(t *testing.T) {
	jsonvalue := func(value any) string {
		v, _ := json.Marshal(value)
		return string(v)
	}

	t.Log(jsonvalue(openapi.ValueToSchema(1)))
	t.Log(jsonvalue(openapi.ValueToSchema("test")))
	t.Log(jsonvalue(openapi.ValueToSchema(struct {
		Test   string
		Bla    int
		Nested []struct {
			Ok bool
		}
	}{})))
}
