package openapi

import (
	"fmt"
)

func requireOpenAPIField(name string, value string) error {
	if len(value) == 0 {
		return fmt.Errorf("OpenAPI Specification v3.1.0: %v is REQUIRED", name)
	}
	return nil
}
