package helpers

import "strings"

// ConvertFlagToKeyValue function convert the key=value list into an interface maps format.
func ConvertFlagToKeyValue(values []string) map[string]interface{} {

	var store = make(map[string]interface{})

	for _, v := range values {
		value := strings.SplitAfter(v, "=")
		store[strings.Trim(strings.ReplaceAll(value[0], "=", ""), " ")] = value[1]
	}

	// store = map[string]interface{}{
	// 	"environments": store,
	// }

	return store
}
