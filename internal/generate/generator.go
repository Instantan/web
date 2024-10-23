package generate

import (
	"bytes"
	"slices"
	"strings"

	"github.com/Instantan/web/internal/openapi"
)

type tsGenerator struct {
	b             *bytes.Buffer
	goalIntent    int
	currentIntent int
}

func newTsGenerator(b *bytes.Buffer) *tsGenerator {
	return &tsGenerator{
		b:             b,
		goalIntent:    0,
		currentIntent: 0,
	}
}

func (t *tsGenerator) intent() *tsGenerator {
	if t.goalIntent != t.currentIntent && (t.goalIntent-t.currentIntent) > 0 {
		must(t.b.WriteString(strings.Repeat("	", t.goalIntent-t.currentIntent)))
		t.currentIntent = t.goalIntent
	}
	return t
}

func (t *tsGenerator) newline() *tsGenerator {
	must(t.b.WriteString("\n"))
	t.currentIntent = 0
	return t
}

func (t *tsGenerator) s(s string) *tsGenerator {
	must(t.b.WriteString(s))
	return t
}

func (t *tsGenerator) name(name string) *tsGenerator {
	t.intent()
	must(t.b.WriteString(name))
	return t
}

func (t *tsGenerator) colon() *tsGenerator {
	must(t.b.WriteString(": "))
	return t
}

func (t *tsGenerator) assign() *tsGenerator {
	must(t.b.WriteString(" = "))
	return t
}

func (t *tsGenerator) scope(b func(t *tsGenerator)) *tsGenerator {
	t.intent()
	must(t.b.WriteString("{"))
	t.newline()
	t.goalIntent++
	b(t)
	t.goalIntent--
	t.newline()
	t.intent()
	must(t.b.WriteString("}"))
	return t
}

func (t *tsGenerator) generic(b func(t *tsGenerator)) *tsGenerator {
	must(t.b.WriteString("<"))
	b(t)
	must(t.b.WriteString(">"))
	return t
}

func (t *tsGenerator) braces(b func(t *tsGenerator)) *tsGenerator {
	t.intent()
	must(t.b.WriteString("("))
	b(t)
	must(t.b.WriteString(")"))
	return t
}

func (t *tsGenerator) semicolon() *tsGenerator {
	must(t.b.WriteString(";"))
	return t
}

func (t *tsGenerator) union() *tsGenerator {
	must(t.b.WriteString(" | "))
	return t
}

func (t *tsGenerator) schema(s openapi.Schema) *tsGenerator {
	t.intent()
	writeSchemaToBuffer(t.b, s, t.goalIntent)
	return t
}

func (t *tsGenerator) doc(lines []string) *tsGenerator {
	must(t.b.WriteString("/**\n"))
	for _, line := range lines {
		must(t.b.WriteString(" *  "))
		must(t.b.WriteString(line))
		must(t.b.WriteString("\n"))
	}
	must(t.b.WriteString("*/"))
	t.newline()
	return t
}

func (t *tsGenerator) comment(s string) *tsGenerator {
	t.intent()
	must(t.b.WriteString("// "))
	must(t.b.WriteString(s))
	t.newline()
	return t
}

func (t *tsGenerator) marker() *tsGenerator {
	t.s("%REMOVEME%")
	return t
}

func (t *tsGenerator) bytes() []byte {
	b := t.b.Bytes()
	b = bytes.ReplaceAll(b, []byte(" | %REMOVEME%"), []byte{})
	b = bytes.ReplaceAll(b, []byte("\n\n%REMOVEME%"), []byte{})
	b = bytes.ReplaceAll(b, []byte("\n%REMOVEME%"), []byte{})
	b = bytes.ReplaceAll(b, []byte("%REMOVEME%"), []byte{})
	return b
}

func writeSchemaToBuffer(b *bytes.Buffer, schema openapi.Schema, indentLevel int) {
	if schema.Ref != "" {
		must(b.WriteString(extractSchemaName(schema.Ref)))
		return
	}
	indent := strings.Repeat("	", indentLevel)
	switch schema.Type {
	case "object":
		if len(schema.Properties) > 0 {
			must(b.WriteString("{\n"))
			for propName, propSchema := range schema.Properties {
				must(b.WriteString(indent + "  " + propName))
				if !slices.Contains(schema.Required, propName) {
					must(b.WriteString("?"))
				}
				must(b.WriteString(": "))
				writeSchemaToBuffer(b, *propSchema, indentLevel+1)
				must(b.WriteString(";\n"))
			}
			b.WriteString(indent + "}")
		} else {
			must(b.WriteString("never"))
		}
	case "array":
		writeSchemaToBuffer(b, *schema.Items, indentLevel)
		must(b.WriteString("[]"))
	case "string":
		must(b.WriteString("string"))
	case "number", "integer":
		must(b.WriteString("number"))
	case "boolean":
		must(b.WriteString("boolean"))
	case "null":
		must(b.WriteString("null"))
	default:
		must(b.WriteString("any"))
	}
}
