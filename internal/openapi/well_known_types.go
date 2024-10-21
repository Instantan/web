package openapi

import (
	"time"
)

func schemaForWellKnownTypes(value any) *Schema {
	switch v := value.(type) {
	case time.Time:
		return timeSchema(v)
	case *time.Time:
		return timeSchema(*v)
	}
	return nil
}

func timeSchema(t time.Time) *Schema {
	iso := t.Format(time.RFC3339)
	return &Schema{
		Type:     "string",
		Example:  iso,
		TypeName: "Time",
	}
}
