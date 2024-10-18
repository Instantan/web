package web

import "fmt"

func assertIsOneOf(value string, allowed []string) {
	for i := range allowed {
		if allowed[i] == value {
			return
		}
	}
	panic(fmt.Errorf("%v must be one of %v", value, allowed))
}

func assertIsNotEmpty(name string, value string) {
	if len(value) == 0 {
		panic(fmt.Errorf("%v must not be EMPTY", name))
	}
}

func assertIsNotNil(name string, value any) {
	if value == nil {
		panic(fmt.Errorf("%v must not be NIL", name))
	}
}
