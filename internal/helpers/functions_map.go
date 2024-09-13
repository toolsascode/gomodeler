package helpers

import (
	"html/template"
	"strings"

	sprig "github.com/Masterminds/sprig/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var funcMapInternal = template.FuncMap{
	"upperCase":        strings.ToUpper,
	"lowerCase":        strings.ToLower,
	"titleCase":        strings.ToTitle,
	"capitalCase":      cases.Title(language.English, cases.Compact).String,
	"pascalCase":       ToPascalCase,
	"camelCase":        ToCamelCase,
	"snakeCaseLower":   ToSnakeCaseLower,
	"snakeCaseUpper":   ToSnakeCaseUpper,
	"kebabCaseLower":   ToKebabCaseLower,
	"kebabCaseUpper":   ToKebabCaseUpper,
	"withoutExtension": WithoutExtension,
	"extension":        Extension,
	"fileName":         GetFileName,
	"dirName":          GetDirectory,
	"sum":              SumFunc,
	"json":             ToJson,
	"yaml":             ToYaml,
	"include":          IncludeFile,
	"count":            strings.Count,
	"compare":          strings.Compare,
	"lastIndex":        strings.LastIndex,
	"contains":         strings.Contains,
	"containsAny":      strings.ContainsAny,
	"hasPrefix":        strings.HasPrefix,
	"hasSuffix":        strings.HasSuffix,
	"split":            strings.Split,
	"splitN":           strings.SplitN,
	"splitAfter":       strings.SplitAfter,
	"splitAfterN":      strings.SplitAfterN,
	"replace":          strings.Replace,
	"replaceAll":       strings.ReplaceAll,
	"join":             strings.Join,
	"bool":             ToBool,
	"arrayStr":         ToArrayString,
}

// Funcs internals and Sprig
// Reference: https://masterminds.github.io/sprig/
func GetFuncMap() template.FuncMap {
	return MergeMaps(funcMapInternal, sprig.GenericFuncMap())
}
