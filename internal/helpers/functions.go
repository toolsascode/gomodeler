package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v3"
)

func ToPascalCase(t string) string {
	title := strings.ReplaceAll(t, "_", " ")
	title = cases.Title(language.English, cases.Compact).String(title)
	return strings.ReplaceAll(title, " ", "")
}

// ToKebabCaseUpper Screaming kebab case
func ToKebabCaseUpper(t string) string {
	upper := strings.ReplaceAll(t, "_", " ")
	upper = cases.Upper(language.English, cases.Compact).String(upper)
	return strings.ReplaceAll(upper, " ", "-")
}

// ToKebabCaseLower Kebab Case
func ToKebabCaseLower(t string) string {
	lower := strings.ReplaceAll(t, "_", " ")
	lower = cases.Lower(language.English, cases.Compact).String(lower)
	return strings.ReplaceAll(lower, " ", "-")
}

func ToSnakeCaseUpper(t string) string {
	upper := strings.ReplaceAll(t, "_", " ")
	upper = cases.Upper(language.English, cases.Compact).String(upper)
	return strings.ReplaceAll(upper, " ", "_")
}

func ToSnakeCaseLower(t string) string {
	lower := strings.ReplaceAll(t, "_", " ")
	lower = cases.Lower(language.English, cases.Compact).String(lower)
	return strings.ReplaceAll(lower, " ", "_")
}

func ToCamelCase(t string) string {
	var group []string
	var firstCase bool

	for _, v := range strings.Fields(t) {
		if !firstCase {
			firstCase = true
			group = append(group, cases.Lower(language.English, cases.Compact).String(v))
		} else {
			group = append(group, cases.Title(language.English, cases.Compact).String(v))
		}
	}

	return strings.Join(group, "")
}

func ToArrayString(input string, character ...string) []string {

	var char string

	if len(character) < 1 {
		char = " "
	} else {
		char = character[0]
	}

	return strings.Split(input, char)
}

func GetDirectory(filename string) string {
	return filepath.Dir(filename)
}

func GetFileName(filename string) string {
	return filepath.Base(filename)
}

func Extension(filename string) string {
	return filepath.Ext(filename)
}

func WithoutExtension(filename string) string {
	var (
		extension = filepath.Ext(filename)
		file      = filepath.Base(filename)
	)
	return strings.TrimRight(file, extension)
}

// Adds the SUM function for the template file
func SumFunc(c, i int) int {
	return c + i
}

func ToJSON(data any) any {
	output, _ := json.Marshal(data)
	return string(output)
}

func ToYAML(data any) any {
	output, _ := yaml.Marshal(data)
	return string(output)
}

func ToBool(t any, valueTrue ...any) bool {

	var valueT any

	if len(valueTrue) > 0 {
		valueT = valueTrue[0]
	}

	switch t {
	case "on":
		return true
	case "true":
		return true
	case "truth":
		return true
	case "verdade":
		return true
	case "verdadeiro":
		return true
	case "1":
		return true
	case true:
		return true
	case 1:
		return true
	case valueT:
		return true
	default:
		return false
	}

}

// IncludeFile - <filename> the file path is relative to the gomodeler CLI
func IncludeFile(filename string) string {
	file, err := os.ReadFile(filename)
	logrus.Debugf("Include file: %s", filename)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func MergeMaps(maps ...map[string]any) map[string]any {
	m := make(map[string]any)
	for _, mps := range maps {
		for k, v := range mps {
			m[k] = v
		}
	}
	return m
}
