package generate

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func generateFunctionName(method, route string) string {
	// Map HTTP methods to action verbs using http.Method constants
	methodMap := map[string]string{
		http.MethodGet:     "Get",
		http.MethodPost:    "Create",
		http.MethodPut:     "Update",
		http.MethodPatch:   "Update",
		http.MethodDelete:  "Delete",
		http.MethodOptions: "Options",
		http.MethodHead:    "Head",
	}

	// Get the action verb for the method
	action, exists := methodMap[method]
	if !exists {
		// Default to capitalized method name
		action = capitalize(strings.ToLower(method))
	}

	// Regular expression to match placeholders like {name}
	placeholderRegex := regexp.MustCompile(`^\{(.+?)\}$`)

	// Split the route into parts
	parts := strings.Split(route, "/")

	var resources []string
	var parameters []string

	// Keep track of parameter counts to handle duplicates
	paramCount := make(map[string]int)

	for _, part := range parts {
		if part == "" {
			continue
		}
		if matches := placeholderRegex.FindStringSubmatch(part); len(matches) == 2 {
			// It's a parameter placeholder
			paramName := capitalize(matches[1])

			// Check for duplicates and append a number if necessary
			paramCount[paramName]++
			if paramCount[paramName] > 1 {
				paramName = fmt.Sprintf("%s%d", paramName, paramCount[paramName])
			}

			parameters = append(parameters, paramName)
		} else {
			// It's a resource name
			resourceName := capitalize(part)
			resources = append(resources, resourceName)
		}
	}

	// Build the function name
	functionName := action
	for _, resource := range resources {
		functionName += resource
	}

	if len(parameters) > 0 {
		functionName += "By"
		functionName += strings.Join(parameters, "And")
	}

	return functionName
}
